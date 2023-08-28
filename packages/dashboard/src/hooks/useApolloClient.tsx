import { FC, PropsWithChildren, useMemo } from 'react';
import { ApolloProvider } from '@apollo/client/react/context';
import { ApolloClient, InMemoryCache, fromPromise } from '@apollo/client/core';
import { createHttpLink } from '@apollo/client/link/http';
import { split } from '@apollo/client/link/core';
import { onError } from '@apollo/client/link/error';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient as createWSClient } from 'graphql-ws';
import memoize from 'lodash.memoize';
import { setContext } from '@apollo/client/link/context';
import { getMainDefinition } from '@apollo/client/utilities';
import queryString from 'query-string';
import { RefreshTokenDocument, RefreshTokenMutation } from './useApollo';
import getConfig from 'next/config.js';

export const deleteCookie = (name: string) => {
  document.cookie = name + '=; Max-Age=0;';
};

export const getCookie = (name: string) => {
  var nameEQ = name + '=';
  var ca = document.cookie.split(';');
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') c = c.substring(1, c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
};

export const setCookie = (name: string, value: string, days: number) => {
  var expires = '';
  if (days) {
    var date = new Date();
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
    expires = '; expires=' + date.toUTCString();
  }
  document.cookie = name + '=' + (value || '') + expires + '; path=/';
};

type GetCookie = typeof getCookie;
type SetCookie = typeof setCookie;
type DeleteCookie = typeof deleteCookie;

const { publicRuntimeConfig: config } = getConfig();

const tokenClient = new ApolloClient({
  uri: config.api.url,
  cache: new InMemoryCache(),
});

const createLink = ({
  getCookie,
  setCookie,
  deleteCookie,
}: {
  getCookie: GetCookie;
  setCookie: SetCookie;
  deleteCookie: DeleteCookie;
}) => {
  let link = createHttpLink({
    uri: config.api.url,
  });

  const refreshToken = memoize(
    async (refreshToken: string | null, propagateError: Error) => {
      if (!refreshToken) {
        if (typeof window !== 'undefined') {
          deleteCookie('accessToken');
          deleteCookie('refreshToken');
          deleteCookie('loginToken');
        }
        throw propagateError;
      }

      const { data: result } = await tokenClient.mutate<RefreshTokenMutation>({
        mutation: RefreshTokenDocument,
        variables: {
          refreshToken,
        },
      });

      setCookie('accessToken', result!.refreshToken.accessToken, 365);
      setCookie('refreshToken', result!.refreshToken.refreshToken, 365);

      return result!.refreshToken.accessToken;
    }
  );

  link = onError(({ graphQLErrors, operation, forward }) => {
    if (graphQLErrors) {
      for (let err of graphQLErrors) {
        switch ((err.extensions?.code as string)?.toUpperCase?.()) {
          case 'UNAUTHENTICATED':
            return fromPromise(
              refreshToken(getCookie('refreshToken'), err)
            ).flatMap(accessToken => {
              const { headers } = operation.getContext();
              operation.setContext({
                headers: {
                  ...headers,
                  authorization: `Bearer ${accessToken}`,
                },
              });

              return forward(operation);
            });
        }
      }
    }
  }).concat(link);

  link = setContext((_, { headers }) => {
    const accessToken = getCookie('accessToken');
    const loginToken = getCookie('loginToken');

    if (!accessToken && !loginToken) {
      return { headers };
    }

    return {
      headers: {
        ...headers,
        authorization: accessToken
          ? `Bearer ${accessToken}`
          : loginToken
          ? `Bearer ${loginToken}`
          : undefined,
      },
    };
  }).concat(link);

  if (typeof window !== 'undefined') {
    const wsLink = new GraphQLWsLink(
      createWSClient({
        retryAttempts: Infinity,
        shouldRetry: () => true,
        keepAlive: 10000,
        url: () => {
          const accessToken = getCookie('accessToken');
          const loginToken = getCookie('loginToken');
          const url = new URL(config.api.url);

          url.protocol =
            url.protocol === 'https:'
              ? 'wss:'
              : url.protocol === 'http:'
              ? 'ws:'
              : url.protocol;

          if (accessToken || loginToken) {
            url.search =
              '?' +
              queryString.stringify({
                bearer: accessToken ?? loginToken ?? undefined,
              });
          }

          return url.toString();
        },
      })
    );

    link = split(
      ({ query }) => {
        const definition = getMainDefinition(query);
        return (
          definition.kind === 'OperationDefinition' &&
          definition.operation === 'subscription'
        );
      },
      wsLink,
      link
    );
  }

  return link;
};

export const createClient = ({
  getCookie,
  setCookie,
  deleteCookie,
}: {
  getCookie: GetCookie;
  setCookie: SetCookie;
  deleteCookie: DeleteCookie;
}) =>
  new ApolloClient({
    link: createLink({ getCookie, setCookie, deleteCookie }),
    cache: new InMemoryCache(),
  });

const withApollo: FC<PropsWithChildren> = ({ children }) => {
  const client = useMemo(
    () => createClient({ getCookie, setCookie, deleteCookie }),
    []
  );
  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};

export default withApollo;

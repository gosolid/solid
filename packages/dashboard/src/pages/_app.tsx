import type { AppProps } from 'next/app';
import withTheme from 'theme';
import ApolloProvider from 'hooks/useApolloClient';
import logo from 'components/Logo';
import 'theme/theme.global.scss';
import 'layouts/site.global.scss';
import { FC, PropsWithChildren } from 'react';

export type Page<P extends any = {}> = FC<P> & {
  getLayout?: FC<PropsWithChildren<P>>;
};

const Theme = withTheme({
  logo,
});

const App = ({ Component, pageProps }: AppProps) => {
  const Layout: FC<PropsWithChildren> =
    (Component as any).getLayout ?? (({ children }) => <>{children}</>);

  return (
    <Theme>
      <ApolloProvider>
        <Layout {...pageProps}>
          <Component {...pageProps} />
        </Layout>
      </ApolloProvider>
    </Theme>
  );
};

export default App;

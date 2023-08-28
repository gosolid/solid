import {
  Children,
  FC,
  ReactElement,
  ReactNode,
  Suspense,
  cloneElement,
  useCallback,
  useMemo,
  useRef,
  useState,
} from 'react';
import { EventEmitter } from 'events';
import { useIsomorphicEffect } from '../hooks/useIsomorphicEffect';
import { createResolver } from '@grexie/resolvable';

type Props = { children: React.ReactNode };

const cloneElementDeep = (element: ReactNode) => {
  return Children.map(element, child => {
    if (typeof child !== 'object' || child === null) {
      return child;
    }
    const props = { children: null as ReactNode };
    if ((child as any).props) {
      props.children = cloneElementDeep((child as any).props.children);
    }
    return cloneElement(child as ReactElement, props);
  });
};

export function useTunnel() {
  const childrenRef = useRef<FC<{}>>(() => null);
  const events = useMemo(() => new EventEmitter(), []);
  const resolver = useMemo(() => createResolver(), []);

  const useChildren = useCallback(() => {
    const [, setState] = useState({});
    useMemo(() => {
      const onUpdate = () => {
        setState({});
      };
      events.on('update', onUpdate);
      return () => {
        events.removeListener('update', onUpdate);
      };
    }, []);

    return childrenRef.current;
  }, []);

  return {
    In: ({ children }: Props) => {
      if (typeof window === 'undefined') {
        // useMemo(() => {
        //   resolver.resolve();
        //   childrenRef.current = () => <>{children}</>;
        //   setImmediate(() => events.emit('update'));
        // }, [children]);
      } else {
        useIsomorphicEffect(() => {
          childrenRef.current = () => <>{children}</>;
          resolver.resolve();
          setImmediate(() => events.emit('update'));

          return () => {
            childrenRef.current = () => null;
            events.emit('update');
          };
        }, [children]);
      }

      return null;
    },

    Out: () => {
      const Children = useChildren() ?? (() => null);

      return <Children />;
    },
  };
}

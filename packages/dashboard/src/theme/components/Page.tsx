import { FC, PropsWithChildren, useState } from 'react';
import { useIsomorphicEffect } from '../hooks/useIsomorphicEffect';
import styles from './Page.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

import { Open_Sans } from 'next/font/google';

export const openSans = Open_Sans({
  display: 'block',
  subsets: ['latin'],
  variable: '--font-primary',
});

export const Page: FC<PropsWithChildren> = ({ children }) => {
  const [javaScript, setJavaScript] = useState(false);

  useIsomorphicEffect(() => {
    setJavaScript(true);
  }, []);

  useIsomorphicEffect(() => {
    const handler = () => {};
    document.addEventListener('touchstart', handler, true);
    return () => document.removeEventListener('touchstart', handler, true);
  }, []);

  return (
    <>
      <div className={cx('root', openSans.variable, { noscript: !javaScript })}>
        {children}
      </div>
    </>
  );
};

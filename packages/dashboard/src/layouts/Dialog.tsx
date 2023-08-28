import { FC, PropsWithChildren, useEffect, useState } from 'react';
import { Logo } from 'theme';
import styles from './dialog.module.scss';
import classNames from 'classnames/bind';
import { useIsomorphicEffect } from 'hooks/useIsomorphicEffect';

const cx = classNames.bind(styles);

export const Dialog: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className={cx('page')}>
      <div className={cx('container')}>
        <header>
          <Logo />
        </header>
        <div className={cx('content')}>{children}</div>
      </div>
    </div>
  );
};

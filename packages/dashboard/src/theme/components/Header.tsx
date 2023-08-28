import {
  ComponentType,
  createElement,
  ElementType,
  forwardRef,
  PropsWithChildren,
} from 'react';
import styles from './Header.module.scss';
import { Logo } from './Logo';
import { Nav } from './Nav';
import { Spacer } from './Spacer';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface HeaderProps<T extends ElementType> {
  component?: string | ComponentType<Omit<HeaderProps<any>, 'component'>>;
  className: string;
}

export const Header = forwardRef<
  HTMLInputElement,
  PropsWithChildren<HeaderProps<any>>
>(({ component = 'header', children, className, ...props }, ref) => {
  return createElement(
    component,
    {
      className: cx(className, 'header'),
    },
    <>
      <Logo className={cx('logo')} />
      <Spacer />
      <Nav auto>{children}</Nav>
    </>
  );
});

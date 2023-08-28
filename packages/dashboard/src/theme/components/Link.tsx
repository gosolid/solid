import {
  createElement,
  ElementType,
  forwardRef,
  PropsWithChildren,
} from 'react';
import { useButton, ButtonProps as ButtonUnstyledProps } from '@mui/base';
import styles from './Link.module.scss';
import classNames from 'classnames/bind';
import { useRouter } from 'next/router';

const cx = classNames.bind(styles);

export type LinkProps<
  T extends ElementType,
  P extends any = {}
> = ButtonUnstyledProps<T> &
  P & {
    primary?: boolean;
    secondary?: boolean;
    disabled?: boolean;
    href?: string;
    replace?: boolean;
  };

export const Link = forwardRef<any, LinkProps<any, any>>(
  (
    {
      component = 'a',
      primary = false,
      secondary = false,
      disabled = false,
      href,
      replace = false,
      children,
      className,
      ...props
    },
    ref
  ) => {
    const { active, focusVisible, getRootProps } = useButton(props);
    const router = useRouter();

    const onClick = (event: MouseEvent) => {
      event.preventDefault();
      router.push(href);
    };

    return createElement(
      component,
      {
        ...getRootProps(),
        href,
        onClick,
        tabIndex: 0,
        className: cx(className, 'link', {
          active,
          disabled,
          focus: focusVisible,
          primary,
          secondary,
        }),
      },
      children
    );
  }
);

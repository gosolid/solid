import {
  createElement,
  ElementType,
  FC,
  forwardRef,
  PropsWithChildren,
} from 'react';
import { useButton, ButtonProps as ButtonUnstyledProps } from '@mui/base';
import styles from './index.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export type ButtonProps<
  T extends ElementType = 'button',
  P extends any = {}
> = ButtonUnstyledProps<T> &
  P & {
    component?: ElementType;
    disabled?: boolean;
    primary?: boolean;
    secondary?: boolean;
    value?: boolean;
  };

export const Button: FC<PropsWithChildren<ButtonProps>> = forwardRef<
  any,
  ButtonProps
>(
  (
    {
      component = 'button',
      primary,
      secondary,
      disabled = false,
      children,
      value,
      className,
      ...props
    },
    ref
  ) => {
    const { active, focusVisible, getRootProps } = useButton(props);

    const el = createElement(
      component,
      {
        ...getRootProps(),
        className: cx(className, 'button', {
          active,
          disabled,
          focus: focusVisible,
          primary,
          secondary,
          value,
        }),
        ref,
      },
      children
    );

    return el;
  }
);

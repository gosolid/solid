import {
  createElement,
  ElementType,
  FC,
  FocusEventHandler,
  PropsWithChildren,
} from 'react';
import styles from './Control.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface ControlProps {
  component?: ElementType;
  className?: string;
  onFocus?: FocusEventHandler;
  onBlur?: FocusEventHandler;
}

export const Control: FC<PropsWithChildren<ControlProps>> = ({
  component = 'div',
  className,
  children,
  onFocus,
  onBlur,
}) => {
  return createElement(
    component,
    {
      className: cx(className, 'control'),
      onFocus,
      onBlur,
    },
    <div className={cx('wrapper')}>{children}</div>
  );
};

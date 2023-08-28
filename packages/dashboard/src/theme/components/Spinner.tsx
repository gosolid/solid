import { FC, HTMLAttributes, useEffect, useState } from 'react';
import styles from './Spinner.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface SpinnerProps {
  className?: string;
  show?: boolean;
  delay?: number;
  align?: 'right' | 'left';
  background?: boolean;
  style?: HTMLAttributes<HTMLDivElement>['style'];
}

export const Spinner: FC<SpinnerProps> = ({
  className,
  delay = 500,
  show: _show = true,
  align,
  style,
  background,
}) => {
  const [show, setShow] = useState<boolean>(false);

  useEffect(() => {
    const timeout = setTimeout(() => setShow(_show), delay);
    return () => clearTimeout(timeout);
  }, [delay, _show]);

  return (
    <div
      className={cx('container', align, { show, background }, className)}
      style={style}
    >
      <div className={cx('spinner')}>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
      </div>
    </div>
  );
};

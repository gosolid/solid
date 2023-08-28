import { FC } from 'react';
import ChevronIcon from './Chevron.svg?icon';
import styles from './Chevron.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface ChevronProps {
  className?: string;
  size?: string | number;
  up?: boolean;
  down?: boolean;
  left?: boolean;
  right?: boolean;
}

export const Chevron: FC<ChevronProps> = ({
  className,
  size,
  up,
  down,
  left,
  right,
}) => {
  return (
    <ChevronIcon
      width={size}
      className={cx({ up, down, left, right }, className)}
    />
  );
};

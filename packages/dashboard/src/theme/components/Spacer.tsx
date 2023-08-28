import { FC } from 'react';
import styles from './Spacer.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface SpacerProps {
  className?: string;
}

export const Spacer: FC<SpacerProps> = ({ className }) => {
  return <div className={cx(className, 'spacer')} />;
};

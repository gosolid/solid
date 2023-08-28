import { ComponentType, FC, createContext, useContext } from 'react';
import styles from './Logo.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface IconProps {
  width: string;
  className: string;
}

export interface LogoContext {
  icon?: ComponentType<IconProps>;
  wordmark?: ComponentType<{}>;
}

export interface LogoProps extends LogoContext {
  className?: string;
}

export const LogoContext = createContext<LogoContext>({});

export const Logo: FC<LogoProps> = ({
  className,
  icon: Icon,
  wordmark: Wordmark,
}) => {
  if (!Icon) {
    Icon = useContext(LogoContext).icon;
  }

  if (!Wordmark) {
    Wordmark = useContext(LogoContext).wordmark;
  }

  return (
    <a className={cx(className, 'logo')} href="/">
      {Icon && <Icon width="1.5em" className={cx('image')} />}
      {Wordmark && (
        <div className={cx('wordmark')}>
          <Wordmark />
        </div>
      )}
    </a>
  );
};

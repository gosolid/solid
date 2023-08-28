import {
  Children,
  cloneElement,
  FC,
  PropsWithChildren,
  ReactElement,
  useEffect,
  useRef,
  useState,
} from 'react';
import { useIsomorphicEffect } from '../hooks/useIsomorphicEffect';
import { Button } from './Button/index';
import styles from './Nav.module.scss';
import buttonStyles from './Button/index.module.scss';
import { Sidebar } from './Sidebar';
import Chevron from './Chevron.svg?icon';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface NavProps {
  className?: string;
  menu?: ReactElement;
  auto?: boolean;
}

export const Nav: FC<PropsWithChildren<NavProps>> = ({
  children,
  className,
  menu: menuLabel = 'Menu',
  auto = false,
  ...props
}) => {
  const [initialized, setInitialized] = useState(false);
  const [menu, setMenu] = useState(!auto);
  const [show, setShow] = useState(false);
  const ref = useRef<HTMLElement | null>(null);
  const button = useRef<HTMLButtonElement | null>(null);
  const menuRef = useRef<HTMLElement | null>(null);
  const [bounds, setBounds] = useState<DOMRect>();

  useIsomorphicEffect(() => setInitialized(true), []);

  useIsomorphicEffect(() => {
    if (!ref.current) {
      return;
    }

    const onResize = () => {
      if (!ref.current) {
        return;
      }

      if (!auto) {
        return;
      }

      const bounds = ref.current.getBoundingClientRect();
      const containerBounds =
        ref.current.parentElement!.getBoundingClientRect();
      setMenu(containerBounds.width < bounds.width);
      if (containerBounds.width < bounds.width) {
        setShow(false);
      }
    };

    const observer = new ResizeObserver(onResize);
    observer.observe(ref.current.parentElement!);
    return () => observer.disconnect();
  }, [ref.current]);

  const onClickMenu = () => {
    setBounds(button.current!.getBoundingClientRect());
    setShow(show => !show);
  };

  useEffect(() => {
    if (show) {
      document.documentElement.classList.add('noscroll');
    } else {
      document.documentElement.classList.remove('noscroll');
    }
  }, [show]);

  const onClickMenuBackground = (event: MouseEvent) => {
    if (
      !(
        menuRef.current === event.target ||
        menuRef.current?.contains(event.target! as any)
      )
    ) {
      setShow(false);
    }
  };

  return (
    <div className={cx('container', { startup: !initialized })}>
      {auto && (
        <nav ref={ref} className={cx(className, 'nav', { hide: !!menu })}>
          {children}
        </nav>
      )}
      {initialized && (
        <>
          <Button
            className={cx('menu-button', { hide: !menu })}
            secondary
            onClick={onClickMenu}
            ref={button}
          >
            {menuLabel}
            <div className={cx('spacer')} />
            <div>
              <Chevron width="1em" className={cx('chevron')} />
            </div>
          </Button>

          <Sidebar
            className={cx('sidebar', { show })}
            onClick={onClickMenuBackground}
          >
            <nav
              className={cx(className, 'menu', { hide: !!menu })}
              style={{
                right: bounds && window.innerWidth - bounds!.right,
                top: bounds?.bottom,
              }}
              ref={menuRef}
            >
              {Children.map(children, child => cloneElement(child as any))}
            </nav>
          </Sidebar>
        </>
      )}
    </div>
  );
};

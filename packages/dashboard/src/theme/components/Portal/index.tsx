import {
  FC,
  MouseEvent,
  PropsWithChildren,
  Suspense,
  createContext,
  useContext,
  useMemo,
} from 'react';
import { Header as MainHeader } from '../Header';
import styles from './index.module.scss';
import { useTunnel } from '../../hooks/useTunnel';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface PortalProps {
  className?: string;
}

export type PortalBlock = ReturnType<typeof useTunnel>['In'];

export interface PortalContext {
  Header: PortalBlock;
  Nav: PortalBlock;
  Toolbar: PortalBlock;
  Sidebar: PortalBlock;
  Footer: PortalBlock;
}

const PortalContext = createContext<PortalContext | null>(null);

export const Portal: FC<PropsWithChildren<PortalProps>> & {
  [k: string]: FC<PropsWithChildren>;
} = (({ children, className }: PropsWithChildren<PortalProps>) => {
  const header = useTunnel();
  const nav = useTunnel();
  const toolbar = useTunnel();
  const sidebar = useTunnel();
  const footer = useTunnel();

  const context = useMemo(
    () => ({
      Header: header.In,
      Nav: nav.In,
      Toolbar: toolbar.In,
      Sidebar: sidebar.In,
      Footer: footer.In,
    }),
    [header.In, nav.In, toolbar.In, sidebar.In, footer.In]
  );

  const onClickNav = (event: MouseEvent<HTMLAnchorElement>) => {
    const target = event.target as HTMLAnchorElement;
    if (target.nextElementSibling?.tagName === 'UL') {
      event.preventDefault();
      target.focus();
    }
  };

  return (
    <PortalContext.Provider value={context}>
      <div className={cx('container', className)}>
        <MainHeader className={cx('header')}>
          <header.Out />
        </MainHeader>
        <div className={cx('content')}>
          <nav className={cx('nav')} onClick={onClickNav}>
            <nav.Out />
          </nav>
          <div className={cx('main')}>
            <div className={cx('toolbar')}>
              <toolbar.Out />
            </div>
            <main>{children}</main>
          </div>
          <div className={cx('sidebar')}>
            <sidebar.Out />
          </div>
        </div>
        <footer className={cx('footer')}>
          <footer.Out />
        </footer>
      </div>
    </PortalContext.Provider>
  );
}) as any;

export const usePortal = () => useContext(PortalContext);

const Header: FC<PropsWithChildren> = ({ children }) => {
  const { Header } = usePortal() ?? {};

  if (!Header) {
    throw new ReferenceError('PortalHeader could not find outer Portal');
  }

  return <Header>{children}</Header>;
};
Portal.Header = Header;

const Footer: FC<PropsWithChildren> = ({ children }) => {
  const { Footer } = usePortal() ?? {};

  if (!Footer) {
    throw new ReferenceError('PortalHeader could not find outer Portal');
  }

  return <Footer>{children}</Footer>;
};
Portal.Footer = Footer;

const Nav: FC<PropsWithChildren> = ({ children }) => {
  const { Nav } = usePortal() ?? {};

  if (!Nav) {
    throw new ReferenceError('PortalHeader could not find outer Portal');
  }

  return <Nav>{children}</Nav>;
};
Portal.Nav = Nav;

const Toolbar: FC<PropsWithChildren> = ({ children }) => {
  const { Toolbar } = usePortal() ?? {};

  if (!Toolbar) {
    throw new ReferenceError('PortalHeader could not find outer Portal');
  }

  return <Toolbar>{children}</Toolbar>;
};
Portal.Toolbar = Toolbar;

const Sidebar: FC<PropsWithChildren> = ({ children }) => {
  const { Sidebar } = usePortal() ?? {};

  if (!Sidebar) {
    throw new ReferenceError('PortalHeader could not find outer Portal');
  }

  return <Sidebar>{children}</Sidebar>;
};
Portal.Sidebar = Sidebar;

import 'core-js';

import { ComponentType, PropsWithChildren, useContext } from 'react';
import { Page } from './Page';
import { Logo, LogoContext } from './Logo';

export * from './Badge';
export * from './BlogIndex';
export * from './Button/index';
export * from './Chevron';
export * from './Control';
export * from './Dialog';
export * from './Field';
export * from './Form';
export * from './Grid/index';
export * from './Header';
export * from './Input';
export * from './Link';
export * from './Logo';
export * from './Menu';
export * from './Nav';
export * from './Notification';
export * from './Option';
export * from './Page';
export * from './Photo';
export * from './Portal/index';
export * from './Select';
export * from './Sidebar';
export * from './Slider';
export * from './Spacer';
export * from './Spinner';
export * from './Switch';
export * from './TablePagination';
export * from './Tabs';
export * from './YouTube';

export type ThemeVariables = Record<string, string>;

export interface ThemeOptions {
  logo?: LogoContext;
}

const withTheme =
  ({ logo }: ThemeOptions = {}): ComponentType<PropsWithChildren> =>
  ({ children }) => {
    const logoContext = useContext(LogoContext);

    return (
      <LogoContext.Provider value={logo ?? logoContext}>
        <Page>{children}</Page>
      </LogoContext.Provider>
    );
  };

export default withTheme;

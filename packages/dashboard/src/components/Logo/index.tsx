import { FC } from 'react';
import Logo from './Logo.svg';
import { LogoContext } from 'theme';

const Wordmark: FC<{}> = () => {
  return (
    <>
      <span>Grexie</span>
      <span>KeyBee</span>
    </>
  );
};

const logo: LogoContext = { icon: Logo, wordmark: Wordmark };

export default logo;

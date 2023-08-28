import { Html, Head, Main, NextScript } from 'next/document';
import { FC } from 'react';

const Document: FC<{}> = () => {
  return (
    <Html lang="en">
      <Head></Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
};

export default Document;

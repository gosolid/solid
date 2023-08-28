declare module 'next/dynamic';

declare module 'theme/Logo' {
  const Component: any;
  export default Component;
}

declare module '*.module.scss' {
  import { StyleFunction } from '@grexie/pages-runtime-styles';
  const styles: StyleFunction;
  export default styles;
}

declare module '*.global.scss' {
  import { StyleFunction } from '@grexie/pages-runtime-styles';
  const styles: StyleFunction;
  export default styles;
}

declare module '*.svg' {
  const Component: any;
  export default Component;
}

declare module '*.svg?icon' {
  import type { FC, SVGProps } from 'react';
  const Component: FC<SVGProps<SVGSVGElement>>;
  export default Component;
}

declare module '*.mdx?' {
  import { ComponentType, PropsWithChildren } from 'react';
  const Component: ComponentType<PropsWithChildren<{}>>;
  export default Component;
}

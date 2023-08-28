declare module '*.svg' {
  const Component: any;
  export default Component;
}

declare module '*.svg?icon' {
  import type { FC, SVGProps } from 'react';
  const Component: FC<SVGProps<SVGSVGElement>>;
  export default Component;
}

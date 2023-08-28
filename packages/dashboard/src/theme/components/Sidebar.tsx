import { FC, PropsWithChildren, useState } from 'react';
import { useIsomorphicEffect } from '../hooks/useIsomorphicEffect';
import { createPortal } from 'react-dom';

export interface SidebarProps {
  className?: string;
  onClick: (this: HTMLDivElement, event: MouseEvent) => void;
}

export const Sidebar: FC<PropsWithChildren<SidebarProps>> = ({
  className,
  children,
  onClick,
}) => {
  const [container, setContainer] = useState<HTMLDivElement>();

  useIsomorphicEffect(() => {
    const container = document.createElement('div');
    if (className) {
      container.className = className;
    }
    document.body.appendChild(container);
    setContainer(container);

    return () => container.remove();
  }, [className]);

  useIsomorphicEffect(() => {
    if (!onClick) {
      return;
    }

    container?.addEventListener('click', onClick, false);

    return () => {
      container?.removeEventListener('click', onClick, false);
    };
  }, [container, onClick]);

  if (!container) {
    return null;
  }

  return createPortal(<>{children}</>, container);
};

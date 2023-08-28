import { FC, MouseEvent, PropsWithChildren, useMemo, useRef } from 'react';
import { createPortal } from 'react-dom';
import { useIsomorphicEffect } from 'theme/hooks';
import styles from './Dialog.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface DialogProps {
  show: boolean;
  onClickBackground?: () => void;
}

export const Dialog: FC<PropsWithChildren<DialogProps>> = ({
  show,
  onClickBackground,
  children,
}) => {
  const dialog = useRef<HTMLDivElement>();

  useIsomorphicEffect(() => {
    dialog.current = document.createElement('div');
    document.body.appendChild(dialog.current);

    return () => {
      dialog.current!.remove();
    };
  }, []);

  const onClickDialog = (event: MouseEvent<HTMLDivElement>) => {
    event.stopPropagation();
  };

  return dialog.current
    ? createPortal(
        <div className={cx('background', { show })} onClick={onClickBackground}>
          <div className={cx('dialog')} onClick={onClickDialog}>
            {children}
          </div>
        </div>,
        dialog.current
      )
    : null;
};

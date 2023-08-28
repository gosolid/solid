import { useCreateSiteMutation, useGetMySitesQuery } from 'hooks/useApollo';
import Link from 'next/link';
import { FC, PropsWithChildren } from 'react';
import { Button, Portal } from 'theme';
import styles from './index.module.scss';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export interface DashboardProps {
  id?: string;
}

export const Dashboard: FC<PropsWithChildren<DashboardProps>> = ({
  children,
  id,
}) => {
  return (
    <Portal>
      <Portal.Footer>&copy; 2023 Grexie</Portal.Footer>
      {children}
    </Portal>
  );
};

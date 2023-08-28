import styles from './index.module.scss';
import {
  ComponentType,
  FC,
  PropsWithChildren,
  cloneElement,
  createContext,
  useContext,
  useMemo,
  useState,
} from 'react';
import { useTunnel } from '../../hooks/useTunnel';
import { DataContext, RowContext } from '../../hooks/useGrid';
import classNames from 'classnames/bind';

const cx = classNames.bind(styles);

export type RowType = 'thead' | 'tbody' | 'tfoot';

interface GridContext {
  columns: ComponentType[];
  rows: { type: RowType; Component: ComponentType }[];
  update: () => void;
}

const GridContext = createContext<GridContext | null>(null);

export interface GridProps {
  className?: string;
  keyExtractor?: (row: any) => string;
  data: Record<string, any>[];
}

export const Grid: FC<PropsWithChildren<GridProps>> = ({
  data,
  className,
  keyExtractor,
  children,
}) => {
  const [, setState] = useState({});

  const context = useMemo(() => {
    return {
      columns: [],
      rows: [],
      update: () => setState({}),
    } as GridContext;
  }, []);

  return (
    <>
      <GridContext.Provider value={context}>{children}</GridContext.Provider>
      <DataContext.Provider value={data}>
        <div className={cx('grid', className)}>
          <table>
            <thead>
              {context.rows
                .filter(row => row.type === 'thead')
                .map(({ Component }, i) =>
                  cloneElement(
                    <>
                      <Component key={`thead-tr-${i}`} />
                    </>
                  )
                )}
            </thead>
            <tbody>
              {data.map((row, i) => (
                <RowContext.Provider
                  key={`tbody-${keyExtractor?.(row) ?? i}`}
                  value={row}
                >
                  {context.rows
                    .filter(row => row.type === 'tbody')
                    .map(({ Component }) =>
                      cloneElement(
                        <>
                          <Component
                            key={`tbody-tr-${keyExtractor?.(row) ?? i}`}
                          />
                        </>
                      )
                    )}
                </RowContext.Provider>
              ))}
            </tbody>
            <tfoot>
              {context.rows
                .filter(row => row.type === 'tfoot')
                .map(({ Component }, i) =>
                  cloneElement(
                    <>
                      <Component key={`tfoot-tr-${i}`} />
                    </>
                  )
                )}
            </tfoot>
          </table>
        </div>
      </DataContext.Provider>
    </>
  );
};

export interface RowProps {
  type?: RowType;
}

export const Row: FC<PropsWithChildren<RowProps>> = ({
  children,
  type = 'tbody',
}) => {
  const grid = useContext(GridContext)!;
  const t = useTunnel();

  useMemo(() => {
    grid.rows.push({ type, Component: t.Out });
    grid.update();
  }, []);

  return <t.In>{children}</t.In>;
};

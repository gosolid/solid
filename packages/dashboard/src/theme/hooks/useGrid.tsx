import { createContext, useContext } from 'react';

export const DataContext = createContext<any[] | null>(null);
export const useData = <T extends any[]>(): T[] => useContext(DataContext)!;
export const RowContext = createContext<any | null>(null);
export const useRow = <T extends any>(): T => useContext(RowContext)!;

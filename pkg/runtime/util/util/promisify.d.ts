/**
 * @moduleName util
@packageDescription
@module util*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 3 @column 0 */
declare module "util" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export interface CustomPromisifyLegacy<TCustom extends Function> extends Function { __promisify__: TCustom; }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export interface CustomPromisifySymbol<TCustom extends Function> extends Function { [promisify.custom]: TCustom; }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export type CustomPromisify<TCustom extends Function> = CustomPromisifySymbol<TCustom> | CustomPromisifyLegacy<TCustom>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<TCustom extends Function>(fn: CustomPromisify<TCustom>): TCustom;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<TResult>(fn: (callback: (err: any, result: TResult) => void) => void): () => Promise<TResult>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify(fn: (callback: (err?: any) => void) => void): () => Promise<void>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, TResult>(fn: (arg1: T1, callback: (err: any, result: TResult) => void) => void): (arg1: T1) => Promise<TResult>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1>(fn: (arg1: T1, callback: (err?: any) => void) => void): (arg1: T1) => Promise<void>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, TResult>(fn: (arg1: T1, arg2: T2, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2) => Promise<TResult>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2>(fn: (arg1: T1, arg2: T2, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2) => Promise<void>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, T3, TResult>(fn: (arg1: T1, arg2: T2, arg3: T3, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2, arg3: T3) => Promise<TResult>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, T3>(fn: (arg1: T1, arg2: T2, arg3: T3, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2, arg3: T3) => Promise<void>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, T3, T4, TResult>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4) => Promise<TResult>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, T3, T4>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4) => Promise<void>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, T3, T4, T5, TResult>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5) => Promise<TResult>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify<T1, T2, T3, T4, T5>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5) => Promise<void>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export function promisify(fn: Function): Function;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/promisify.go @line 34 @column 0 */

export namespace promisify { const custom: unique symbol; }}

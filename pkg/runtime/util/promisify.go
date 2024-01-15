//js:package util

package util

//go:generate go run github.com/grexie/isolates/codegen

import (
	"context"
	"fmt"

	isolates "github.com/grexie/isolates"
)

//js:method
//js:export promisify
//ts:export interface CustomPromisifyLegacy<TCustom extends Function> extends Function { __promisify__: TCustom; }
//ts:export interface CustomPromisifySymbol<TCustom extends Function> extends Function { [promisify.custom]: TCustom; }
//ts:export type CustomPromisify<TCustom extends Function> = CustomPromisifySymbol<TCustom> | CustomPromisifyLegacy<TCustom>;
//ts:export function promisify<TCustom extends Function>(fn: CustomPromisify<TCustom>): TCustom;
//ts:export function promisify<TResult>(fn: (callback: (err: any, result: TResult) => void) => void): () => Promise<TResult>;
//ts:export function promisify(fn: (callback: (err?: any) => void) => void): () => Promise<void>;
//ts:export function promisify<T1, TResult>(fn: (arg1: T1, callback: (err: any, result: TResult) => void) => void): (arg1: T1) => Promise<TResult>;
//ts:export function promisify<T1>(fn: (arg1: T1, callback: (err?: any) => void) => void): (arg1: T1) => Promise<void>;
//ts:export function promisify<T1, T2, TResult>(fn: (arg1: T1, arg2: T2, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2) => Promise<TResult>;
//ts:export function promisify<T1, T2>(fn: (arg1: T1, arg2: T2, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2) => Promise<void>;
//ts:export function promisify<T1, T2, T3, TResult>(fn: (arg1: T1, arg2: T2, arg3: T3, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2, arg3: T3) => Promise<TResult>;
//ts:export function promisify<T1, T2, T3>(fn: (arg1: T1, arg2: T2, arg3: T3, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2, arg3: T3) => Promise<void>;
//ts:export function promisify<T1, T2, T3, T4, TResult>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4) => Promise<TResult>;
//ts:export function promisify<T1, T2, T3, T4>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4) => Promise<void>;
//ts:export function promisify<T1, T2, T3, T4, T5, TResult>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5, callback: (err: any, result: TResult) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5) => Promise<TResult>;
//ts:export function promisify<T1, T2, T3, T4, T5>(fn: (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5, callback: (err?: any) => void) => void): (arg1: T1, arg2: T2, arg3: T3, arg4: T4, arg5: T5) => Promise<void>;
//ts:export function promisify(fn: Function): Function;
//ts:export namespace promisify { const custom: unique symbol; }
func Promisify(ctx context.Context, fn *isolates.Value) (isolates.Function, error) {
	if !fn.IsKind(isolates.KindFunction) {
		return nil, fmt.Errorf("util.promisify: not a function %s", fn)
	}

	return func(in isolates.FunctionArgs) (*isolates.Value, error) {
		if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
			return nil, err
		} else if callback, err := in.Context.Create(in.ExecutionContext, func(in isolates.FunctionArgs) (*isolates.Value, error) {
			if in.Arg(in.ExecutionContext, 0).IsNil() {
				return nil, resolver.Resolve(in.ExecutionContext, in.Arg(in.ExecutionContext, 1))
			} else {
				return nil, resolver.Reject(in.ExecutionContext, in.Arg(in.ExecutionContext, 0))
			}
		}); err != nil {
			return nil, err
		} else {
			args := append(in.Args, callback)
			if _, err := fn.CallValue(in.ExecutionContext, in.This, args...); err != nil {
				return nil, err
			} else {
				return resolver.Promise(in.ExecutionContext)
			}
		}
	}, nil
}

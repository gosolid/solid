/**
 * @moduleName stream
@packageDescription
@module stream*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 3 @column 0 */
declare module "stream" {
  import type { Stream } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
  class Readable extends Stream {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 258 @column 0 */
  pause(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 263 @column 0 */
  resume(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 274 @column 0 */
  pipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 328 @column 0 */
  readV8(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 382 @column 0 */
  setEncoding(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 387 @column 0 */
  unpipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 429 @column 0 */
  unshift(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 446 @column 0 */
  wrap(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 451 @column 0 */
  compose(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 456 @column 0 */
  iterator(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 461 @column 0 */
  map(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 466 @column 0 */
  filter(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 471 @column 0 */
  forEach(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 476 @column 0 */
  toArray(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 481 @column 0 */
  some(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 486 @column 0 */
  find(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 491 @column 0 */
  every(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 496 @column 0 */
  flatMap(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 501 @column 0 */
  drop(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 506 @column 0 */
  take(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 511 @column 0 */
  asIndexedPairs(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 516 @column 0 */
  reduce(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 521 @column 0 */
  push(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly isPaused: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readable: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableAborted: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableDidRead: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableEncoding: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableFlowing: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableHighWaterMark: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableLength: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 109 @column 0 */
    readonly readableObjectMode: boolean;
  }
  export { Readable };
}

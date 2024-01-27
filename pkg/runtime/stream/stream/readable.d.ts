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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
  class Readable extends Stream {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 259 @column 0 */
  pause(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 264 @column 0 */
  resume(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 275 @column 0 */
  read(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 293 @column 0 */
  pipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 403 @column 0 */
  setEncoding(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 408 @column 0 */
  unpipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 450 @column 0 */
  unshift(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 467 @column 0 */
  wrap(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 472 @column 0 */
  compose(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 477 @column 0 */
  iterator(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 482 @column 0 */
  map(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 487 @column 0 */
  filter(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 492 @column 0 */
  forEach(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 497 @column 0 */
  toArray(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 502 @column 0 */
  some(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 507 @column 0 */
  find(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 512 @column 0 */
  every(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 517 @column 0 */
  flatMap(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 522 @column 0 */
  drop(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 527 @column 0 */
  take(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 532 @column 0 */
  asIndexedPairs(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 537 @column 0 */
  reduce(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 542 @column 0 */
  push(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableLength: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableObjectMode: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly isPaused: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readable: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableDidRead: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableAborted: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableEncoding: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableFlowing: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 110 @column 0 */
    readonly readableHighWaterMark: number;
  }
  export { Readable };
}

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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/duplex.go @line 3 @column 0 */
declare module "stream" {
  import type { Stream, Writable, Readable } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/duplex.go @line 33 @column 0 */
  class Duplex extends Stream implements Writable, Readable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 169 @column 0 */
  cork(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 174 @column 0 */
  end(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 286 @column 0 */
  setDefaultEncoding(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 291 @column 0 */
  uncork(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 385 @column 0 */
  write(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableNeedDrain: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writable: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableAborted: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    writableHighWaterMark: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableLength: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableObjectMode: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableCorked: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 77 @column 0 */
    readonly writableFinished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 293 @column 0 */
  pause(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 298 @column 0 */
  resume(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 322 @column 0 */
  read(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 348 @column 0 */
  pipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 489 @column 0 */
  setEncoding(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 494 @column 0 */
  unpipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 536 @column 0 */
  unshift(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 572 @column 0 */
  wrap(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 577 @column 0 */
  compose(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 582 @column 0 */
  iterator(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 587 @column 0 */
  map(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 592 @column 0 */
  filter(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 597 @column 0 */
  forEach(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 602 @column 0 */
  toArray(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 607 @column 0 */
  some(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 612 @column 0 */
  find(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 617 @column 0 */
  every(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 622 @column 0 */
  flatMap(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 627 @column 0 */
  drop(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 632 @column 0 */
  take(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 637 @column 0 */
  asIndexedPairs(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 642 @column 0 */
  reduce(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 731 @column 0 */
  push(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableEncoding: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readableHighWaterMark: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableLength: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly isPaused: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readable: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableFlowing: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableObjectMode: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableAborted: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 118 @column 0 */
    readonly readableDidRead: boolean;
  }
  export { Duplex };
}

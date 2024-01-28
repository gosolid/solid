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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
  class Readable extends Stream {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 290 @column 0 */
  pause(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 295 @column 0 */
  resume(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 319 @column 0 */
  read(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 345 @column 0 */
  pipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 486 @column 0 */
  setEncoding(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 491 @column 0 */
  unpipe(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 533 @column 0 */
  unshift(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 569 @column 0 */
  wrap(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 574 @column 0 */
  compose(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 579 @column 0 */
  iterator(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 584 @column 0 */
  map(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 589 @column 0 */
  filter(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 594 @column 0 */
  forEach(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 599 @column 0 */
  toArray(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 604 @column 0 */
  some(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 609 @column 0 */
  find(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 614 @column 0 */
  every(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 619 @column 0 */
  flatMap(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 624 @column 0 */
  drop(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 629 @column 0 */
  take(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 634 @column 0 */
  asIndexedPairs(  ): Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 639 @column 0 */
  reduce(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 714 @column 0 */
  push(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    set setReadableHighWaterMark(value: number);
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableLength: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableObjectMode: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly isPaused: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableFlowing: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableDidRead: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableEncoding: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableHighWaterMark: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readable: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/readable.go @line 116 @column 0 */
    readonly readableAborted: boolean;
  }
  export { Readable };
}

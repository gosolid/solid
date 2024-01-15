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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 3 @column 0 */
declare module "stream" {
  import type { Stream } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
  class Writable extends Stream {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 139 @column 0 */
  cork(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 144 @column 0 */
  end(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 206 @column 0 */
  setDefaultEncoding(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 211 @column 0 */
  uncork(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 300 @column 0 */
  write(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writable: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableAborted: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableHighWaterMark: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableCorked: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableFinished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableLength: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableNeedDrain: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/writable.go @line 67 @column 0 */
    readonly writableObjectMode: boolean;
  }
  export { Writable };
}

/**
 * @moduleName tty
@packageDescription
@module tty*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 3 @column 0 */
declare module "tty" {
  import type { Writable } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 33 @column 0 */
  class WriteStream extends Writable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 64 @column 0 */
  clear(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 73 @column 0 */
  getWindowSize(  ): number[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 33 @column 0 */
    readonly rows: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 33 @column 0 */
    readonly columns: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/writestream.go @line 33 @column 0 */
    readonly isTTY: boolean;
  }
  export { WriteStream };
}

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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/readstream.go @line 3 @column 0 */
declare module "tty" {
  import type { Readable } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/readstream.go @line 33 @column 0 */
  class ReadStream extends Readable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/readstream.go @line 79 @column 0 */
  setRawMode(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/readstream.go @line 33 @column 0 */
    readonly isRaw: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/tty/readstream.go @line 33 @column 0 */
    readonly isTTY: boolean;
  }
  export { ReadStream };
}

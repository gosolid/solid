/**
 * @moduleName buffer
@packageDescription
@module buffer*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 3 @column 0 */
declare module "buffer" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 31 @column 0 */
  class Buffer {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 106 @column 0 */
  readUInt32BE(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 116 @column 0 */
  writeUInt32BE(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 153 @column 0 */
  slice(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 198 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 31 @column 0 */
    readonly buffer: any;
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 45 @column 0 */
  class Buffer {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 106 @column 0 */
  readUInt32BE(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 116 @column 0 */
  writeUInt32BE(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 153 @column 0 */
  slice(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 198 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 45 @column 0 */
    readonly buffer: any;
  }
  export { Buffer };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/buffer/buffer.go @line 68 @column 0 */
  function createBufferFunction(  ): any;
}

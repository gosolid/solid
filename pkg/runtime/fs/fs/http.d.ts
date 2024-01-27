/**
 * @moduleName fs
@packageDescription
@module fs*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 3 @column 0 */
declare module "fs" {
  import type { Buffer } from "buffer";
  import type { File, ReadStream, FileDescriptor, Stats } from "fs";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 38 @column 0 */
  class HttpFileSystem extends File {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 43 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 58 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 58 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 71 @column 0 */
  readFileSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 71 @column 0 */
  readFile(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 102 @column 0 */
  createReadStream(  ): ReadStream;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 120 @column 0 */
  openSync(  ): FileDescriptor;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 120 @column 0 */
  open(  ): FileDescriptor;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 142 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 142 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 147 @column 0 */
  file(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 153 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 153 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 169 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 169 @column 0 */
  stat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 197 @column 0 */
  lstatSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 197 @column 0 */
  lstat(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 203 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 203 @column 0 */
  realpathSync(  ): string;
  }
  export { HttpFileSystem };
}

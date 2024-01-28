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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 3 @column 0 */
declare module "fs" {
  import type { Buffer } from "buffer";
  import type { Stats } from "fs";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 28 @column 0 */
  class File {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 256 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 256 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 263 @column 0 */
  readAllSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 263 @column 0 */
  readAll(  ): Buffer;
  }
  export { File };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 33 @column 0 */
  class FileSystem {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 42 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 47 @column 0 */
  mount(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 53 @column 0 */
  createReadStream(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 65 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 65 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 77 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 77 @column 0 */
  realpathSync(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 94 @column 0 */
  readFileSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 94 @column 0 */
  readFile(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 134 @column 0 */
  openSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 134 @column 0 */
  open(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 146 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 146 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 151 @column 0 */
  file(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 157 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 157 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 169 @column 0 */
  statSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 169 @column 0 */
  stat(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 181 @column 0 */
  lstatSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 181 @column 0 */
  lstat(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 193 @column 0 */
  accessSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 193 @column 0 */
  access(  ): void;
  }
  export { FileSystem };
}

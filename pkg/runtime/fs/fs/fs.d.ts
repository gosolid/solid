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
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 228 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 228 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 235 @column 0 */
  readAllSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 235 @column 0 */
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
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 106 @column 0 */
  openSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 106 @column 0 */
  open(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 118 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 118 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 123 @column 0 */
  file(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 129 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 129 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 141 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 141 @column 0 */
  stat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 153 @column 0 */
  lstatSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 153 @column 0 */
  lstat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 165 @column 0 */
  accessSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/fs.go @line 165 @column 0 */
  access(  ): void;
  }
  export { FileSystem };
}

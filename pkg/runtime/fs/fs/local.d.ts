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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 3 @column 0 */
declare module "fs" {
  import type { Stats, File, FileDescriptor } from "fs";
  import type { Buffer } from "buffer";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 39 @column 0 */
  class LocalFileSystem extends Stats {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 44 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 57 @column 0 */
  mount(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 64 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 64 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 91 @column 0 */
  readFileSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 91 @column 0 */
  readFile(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 125 @column 0 */
  createReadStream(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 141 @column 0 */
  openSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 141 @column 0 */
  open(  ): FileDescriptor;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 163 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 163 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 179 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 179 @column 0 */
  stat(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 207 @column 0 */
  lstatSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 207 @column 0 */
  lstat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 235 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 235 @column 0 */
  realpathSync(  ): string;
  }
  export { LocalFileSystem };
}

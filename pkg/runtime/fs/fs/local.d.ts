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
  import type { File, Stats } from "fs";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 38 @column 0 */
  class LocalFileSystem extends File {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 43 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 56 @column 0 */
  mount(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 63 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 63 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 90 @column 0 */
  readFileSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 90 @column 0 */
  readFile(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 111 @column 0 */
  createReadStream(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 127 @column 0 */
  openSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 127 @column 0 */
  open(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 149 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 149 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 165 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 165 @column 0 */
  stat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 193 @column 0 */
  lstatSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 193 @column 0 */
  lstat(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 221 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/local.go @line 221 @column 0 */
  realpathSync(  ): string;
  }
  export { LocalFileSystem };
}

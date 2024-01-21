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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 3 @column 0 */
declare module "fs" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
  class Stats {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 32 @column 0 */
  isDirectory(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 37 @column 0 */
  isFile(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 42 @column 0 */
  isSymbolicLink(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly mtimeMs: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly atime: Date;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly ctimeMs: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly birthtimeMs: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly gid: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly mtime: Date;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly ctime: Date;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly uid: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly size: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly mode: Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly contentType: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly atimeMs: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go @line 16 @column 0 */
    readonly birthtime: Date;
  }
  export { Stats };
}

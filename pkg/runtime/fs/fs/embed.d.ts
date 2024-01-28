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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 3 @column 0 */
declare module "fs" {
  import type { Stats, File } from "fs";
  import type { Buffer } from "buffer";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 40 @column 0 */
  class EmbeddedFileSystem extends Stats {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 45 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 50 @column 0 */
  createReadStream(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 66 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 66 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 87 @column 0 */
  readFileSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 87 @column 0 */
  readFile(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 116 @column 0 */
  openSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 116 @column 0 */
  open(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 133 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 133 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 138 @column 0 */
  file(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 144 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 144 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 160 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 160 @column 0 */
  stat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 190 @column 0 */
  lstatSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 190 @column 0 */
  lstat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 202 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 202 @column 0 */
  realpathSync(  ): string;
  }
  export { EmbeddedFileSystem };
}

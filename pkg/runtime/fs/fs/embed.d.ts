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
  import type { Buffer } from "buffer";
  import type { File, ReadStream, FileDescriptor, Stats } from "fs";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 40 @column 0 */
  class EmbeddedFileSystem extends File {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 45 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 50 @column 0 */
  createReadStream(  ): ReadStream;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 66 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 66 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 87 @column 0 */
  readFileSync(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 87 @column 0 */
  readFile(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 103 @column 0 */
  openSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 103 @column 0 */
  open(  ): FileDescriptor;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 120 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 120 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 125 @column 0 */
  file(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 131 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 131 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 147 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 147 @column 0 */
  stat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 177 @column 0 */
  lstatSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 177 @column 0 */
  lstat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 189 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/embed.go @line 189 @column 0 */
  realpathSync(  ): string;
  }
  export { EmbeddedFileSystem };
}

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
  import type { ReadStream, Stats, File } from "fs";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 37 @column 0 */
  class HttpFileSystem extends File {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 42 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 57 @column 0 */
  readdirSync(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 57 @column 0 */
  readdir(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 70 @column 0 */
  readFileSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 70 @column 0 */
  readFile(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 87 @column 0 */
  createReadStream(  ): ReadStream;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 105 @column 0 */
  openSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 105 @column 0 */
  open(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 127 @column 0 */
  closeSync(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 127 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 132 @column 0 */
  file(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 138 @column 0 */
  existsSync(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 138 @column 0 */
  exists(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 154 @column 0 */
  statSync(  ): File;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 154 @column 0 */
  stat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 182 @column 0 */
  lstatSync(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 182 @column 0 */
  lstat(  ): Stats;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 188 @column 0 */
  realpath(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/fs/http.go @line 188 @column 0 */
  realpathSync(  ): string;
  }
  export { HttpFileSystem };
}

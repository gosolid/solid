/**
 * @moduleName console
@packageDescription
@module console*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 3 @column 0 */
declare module "console" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 51 @column 0 */
  class Console {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 108 @column 0 */
  assert(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 113 @column 0 */
  clear(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 127 @column 0 */
  count(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 132 @column 0 */
  countReset(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 137 @column 0 */
  debug(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 150 @column 0 */
  dir(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 155 @column 0 */
  dirXML(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 160 @column 0 */
  error(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 173 @column 0 */
  group(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 178 @column 0 */
  groupCollapsed(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 183 @column 0 */
  groupEnd(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 188 @column 0 */
  info(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 193 @column 0 */
  log(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 206 @column 0 */
  table(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 211 @column 0 */
  time(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 216 @column 0 */
  timeEnd(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 221 @column 0 */
  timeLog(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 226 @column 0 */
  trace(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 239 @column 0 */
  warn(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 244 @column 0 */
  profile(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 249 @column 0 */
  profileEnd(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 254 @column 0 */
  timeStamp(  ): void;
  }
  export { Console };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/console/console.go @line 79 @column 0 */
  function createDefaultConsole(  ): Console;
}

/**
 * @moduleName util
@packageDescription
@module util*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 3 @column 0 */
declare module "util" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 26 @column 0 */
  class Counter {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 37 @column 0 */
  clone(  ): Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 53 @column 0 */
  floorTime(  ): Date;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 85 @column 0 */
  add(  ): Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 95 @column 0 */
  addNow(  ): Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 100 @column 0 */
  addDuration(  ): Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 127 @column 0 */
  addDurationNow(  ): Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 132 @column 0 */
  sub(  ): Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 137 @column 0 */
  sumNow(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 142 @column 0 */
  sum(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 186 @column 0 */
  sumAllNow(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 191 @column 0 */
  sumAll(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 196 @column 0 */
  sumDurationNow(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 201 @column 0 */
  sumAllDurationNow(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 206 @column 0 */
  sumDuration(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/util/counter.go @line 211 @column 0 */
  sumAllDuration(  ): number;
  }
  export { Counter };
}

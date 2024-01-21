/**
 * @moduleName assert
@packageDescription
@module assert*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 3 @column 0 */
declare module "assert" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
  class AssertionError {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 71 @column 0 */
  toString(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly message: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly name: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly actual: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly expected: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly operator: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly code: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/assert/errors.go @line 43 @column 0 */
    readonly generatedMessage: boolean;
  }
  export { AssertionError };
}

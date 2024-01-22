/**
 * @moduleName native:@grexie/workers/crypto
@packageDescription
@module native:@grexie/workers/crypto*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/crypto/crypto.go @line 2 @column 0 */
declare module "native:@grexie/workers/crypto" {
  import type { Buffer } from "buffer";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/crypto/crypto.go @line 29 @column 0 */
  class Hash {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/crypto/crypto.go @line 68 @column 0 */
  update(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/crypto/crypto.go @line 74 @column 0 */
  digest(  ): Buffer;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/crypto/crypto.go @line 29 @column 0 */
    readonly algorithm: string;
  }
  export { Hash };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/crypto/crypto.go @line 80 @column 0 */
  function createHmac(  ): any;
  export { createHmac };
}

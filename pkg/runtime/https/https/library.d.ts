/**
 * @moduleName https
@packageDescription
@module https*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 3 @column 0 */
declare module "https" {
  import type { Socket } from "net";
  import type { AgentBase, Https } from "https";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 26 @column 0 */
  class Https {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 26 @column 0 */
    readonly net: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 26 @column 0 */
    readonly agent: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 26 @column 0 */
    readonly globalAgent: AgentBase;
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 34 @column 0 */
  class Https {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 34 @column 0 */
    readonly net: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 34 @column 0 */
    readonly agent: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 34 @column 0 */
    readonly globalAgent: Https;
  }
  export { Https };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/https/library.go @line 78 @column 0 */
  function createDefaultHttp(  ): Https;
}

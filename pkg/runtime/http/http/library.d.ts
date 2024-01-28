/**
 * @moduleName http
@packageDescription
@module http*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 3 @column 0 */
declare module "http" {
  import type { Socket } from "net";
  import type { AgentBase } from "http";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 30 @column 0 */
  class Http {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 30 @column 0 */
    readonly net: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 30 @column 0 */
    readonly server: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 30 @column 0 */
    readonly agent: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 30 @column 0 */
    readonly globalAgent: AgentBase;
  }
  export { Http };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/library.go @line 90 @column 0 */
  function createDefaultHttp(  ): Http;
}

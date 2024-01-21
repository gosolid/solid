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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/agent.go @line 2 @column 0 */
declare module "http" {
  import type { Socket } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/agent.go @line 25 @column 0 */
  class AgentBase {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/agent.go @line 56 @column 0 */
  createConnection(  ): Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/agent.go @line 73 @column 0 */
  getName(  ): string;
  }
}

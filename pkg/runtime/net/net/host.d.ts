/**
 * @moduleName net
@packageDescription
@module net*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 3 @column 0 */
declare module "net" {
  import type { Socket, Net } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 64 @column 0 */
  class HostNetOptions {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 28 @column 0 */
  listen(  ): Socket;
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 74 @column 0 */
  function createHostNet(  ): Net;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 85 @column 0 */
  class HostSocketOptions {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 107 @column 0 */
  localAddress(  ): Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 116 @column 0 */
  remoteAddress(  ): Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 125 @column 0 */
  final(  ): void;
  }
}

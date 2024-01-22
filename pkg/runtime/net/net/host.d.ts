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
  import type { Socket, SocketAddress } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 159 @column 0 */
  class HostNetOptions {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 93 @column 0 */
  connect(  ): Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 123 @column 0 */
  listen(  ): Socket;
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 170 @column 0 */
  function createHostNet(  ): Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 181 @column 0 */
  class HostSocketOptions {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 203 @column 0 */
  localAddress(  ): Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 212 @column 0 */
  remoteAddress(  ): SocketAddress;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/host.go @line 221 @column 0 */
  final(  ): void;
  }
}

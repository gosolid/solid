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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 2 @column 0 */
declare module "net" {
  import type { Duplex } from "stream";
  import type { SocketAddress } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
  class Socket extends Duplex {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 57 @column 0 */
  address(  ): SocketAddress;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 116 @column 0 */
  setTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 121 @column 0 */
  setNoDelay(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
    readonly localAddress: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
    readonly localPort: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
    readonly localFamily: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
    readonly remoteAddress: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
    readonly remotePort: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/net/socket.go @line 19 @column 0 */
    readonly remoteFamily: string;
  }
}

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

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 3 @column 0 */
declare module "http" {
  import type { Readable, Writable } from "stream";
  import type { EventEmitter } from "events";
  import type { Socket } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
  class IncomingMessage extends Readable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 500 @column 0 */
  setTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly httpVersion: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly rawHeaders: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly rawTrailers: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailers: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly complete: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly headers: Record<string, any>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly headersDistinct: Record<string, string[]>;
  }
  export { IncomingMessage };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
  class ServerRequest extends EventEmitter implements IncomingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
    readonly method: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
    readonly url: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 500 @column 0 */
  setTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly headersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly httpVersion: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly rawHeaders: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly rawTrailers: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailers: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly complete: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly headers: Record<string, any>;
  }
  export { ServerRequest };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 282 @column 0 */
  class ClientResponse extends EventEmitter implements IncomingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 282 @column 0 */
    readonly statusMessage: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 282 @column 0 */
    readonly statusCode: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 500 @column 0 */
  setTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailers: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly complete: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly headers: Record<string, any>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly headersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly httpVersion: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly rawHeaders: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly rawTrailers: string[];
  }
  export { ClientResponse };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
  class OutgoingMessage extends Writable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 533 @column 0 */
  addTrailers(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 542 @column 0 */
  appendHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 559 @column 0 */
  getHeader(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 564 @column 0 */
  getHeaderNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 575 @column 0 */
  getRawHeaderNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 591 @column 0 */
  hasHeader(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 601 @column 0 */
  removeHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 606 @column 0 */
  setHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 639 @column 0 */
  setHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 650 @column 0 */

setTimeout(timeout: number, listener: () => void): void  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 685 @column 0 */
  writeContinue(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 690 @column 0 */
  writeEarlyHints(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 695 @column 0 */
  writeProcessing(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 700 @column 0 */
  wait(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly writableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly writableFinished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    sendDate: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    strictContentLength: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly destroyed: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly finished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly headers: Record<string, any>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly headersSent: boolean;
  }
  export { OutgoingMessage };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 404 @column 0 */
  class ServerResponse extends OutgoingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 768 @column 0 */
  flushHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 796 @column 0 */
  writeHead(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 404 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 404 @column 0 */
    statusCode: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 404 @column 0 */
    statusMessage: string;
  }
  export { ServerResponse };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
  class ClientRequest extends OutgoingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 837 @column 0 */
  flushHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 913 @column 0 */
  setNoDelay(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 928 @column 0 */
  setSocketKeepAlive(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    host: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    maxHeadersCount: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    method: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    path: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    protocol: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    readonly reusedSocket: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 430 @column 0 */
    readonly socket: Socket;
  }
  export { ClientRequest };
}

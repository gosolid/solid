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
  import type { Socket } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
  class IncomingMessage extends Readable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 273 @column 0 */
  setTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly headers: Record<string, any>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly headersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly httpVersion: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly rawHeaders: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly rawTrailers: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly trailers: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly trailersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 89 @column 0 */
    readonly complete: boolean;
  }
  export { IncomingMessage };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 104 @column 0 */
  class ServerRequest extends IncomingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 104 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 104 @column 0 */
    readonly method: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 104 @column 0 */
    readonly url: string;
  }
  export { ServerRequest };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
  class OutgoingMessage extends Writable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 306 @column 0 */
  addTrailers(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 315 @column 0 */
  appendHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 332 @column 0 */
  getHeader(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 337 @column 0 */
  getHeaderNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 348 @column 0 */
  getRawHeaderNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 364 @column 0 */
  hasHeader(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 374 @column 0 */
  removeHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 379 @column 0 */
  setHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 412 @column 0 */
  setHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 423 @column 0 */

setTimeout(timeout: number, listener: () => void): void  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 458 @column 0 */
  writeContinue(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 463 @column 0 */
  writeEarlyHints(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 468 @column 0 */
  writeProcessing(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 473 @column 0 */
  wait(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    readonly writableFinished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    sendDate: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    strictContentLength: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    readonly destroyed: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    readonly finished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    readonly headers: Record<string, any>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    readonly headersSent: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 155 @column 0 */
    readonly writableEnded: boolean;
  }
  export { OutgoingMessage };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 208 @column 0 */
  class ServerResponse extends OutgoingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 527 @column 0 */
  flushHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 555 @column 0 */
  writeHead(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 208 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 208 @column 0 */
    statusCode: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 208 @column 0 */
    statusMessage: string;
  }
  export { ServerResponse };
}

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
  import type { Writable, Readable } from "stream";
  import type { EventEmitter } from "events";
  import type { Socket } from "net";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
  class IncomingMessage extends Readable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 498 @column 0 */
  setTimeout(  ): void;
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
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailers: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly trailersDistinct: Record<string, string[]>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 114 @column 0 */
    readonly complete: boolean;
  }
  export { IncomingMessage };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
  class ServerRequest extends IncomingMessage implements EventEmitter {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
    readonly method: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 246 @column 0 */
    readonly url: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 62 @column 0 */
  on(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 62 @column 0 */
  on(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 84 @column 0 */
  once(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 106 @column 0 */
  emit(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 154 @column 0 */
  eventNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 172 @column 0 */
  getMaxListeners(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 177 @column 0 */
  setMaxListeners(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 182 @column 0 */
  listenerCount(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 209 @column 0 */
  rawListeners(  ): any[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 209 @column 0 */
  rawListeners(  ): any[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 231 @column 0 */
  off(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 231 @column 0 */
  off(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 264 @column 0 */
  prependListener(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 286 @column 0 */
  prependListenerOnce(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 308 @column 0 */
  removeAllListeners(  ): void;
  }
  export { ServerRequest };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 283 @column 0 */
  class ClientResponse extends IncomingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 283 @column 0 */
    readonly statusCode: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 283 @column 0 */
    readonly statusMessage: string;
  }
  export { ClientResponse };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
  class OutgoingMessage extends Writable {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 531 @column 0 */
  addTrailers(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 540 @column 0 */
  appendHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 557 @column 0 */
  getHeader(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 562 @column 0 */
  getHeaderNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 573 @column 0 */
  getRawHeaderNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 589 @column 0 */
  hasHeader(  ): boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 599 @column 0 */
  removeHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 604 @column 0 */
  setHeader(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 637 @column 0 */
  setHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 648 @column 0 */

setTimeout(timeout: number, listener: () => void): void  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 683 @column 0 */
  writeContinue(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 688 @column 0 */
  writeEarlyHints(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 693 @column 0 */
  writeProcessing(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 698 @column 0 */
  wait(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly destroyed: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly finished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly headers: Record<string, any>;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly headersSent: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly writableEnded: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    readonly writableFinished: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    sendDate: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 345 @column 0 */
    strictContentLength: boolean;
  }
  export { OutgoingMessage };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 402 @column 0 */
  class ServerResponse extends OutgoingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 766 @column 0 */
  flushHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 794 @column 0 */
  writeHead(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 402 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 402 @column 0 */
    statusCode: number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 402 @column 0 */
    statusMessage: string;
  }
  export { ServerResponse };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
  class ClientRequest extends OutgoingMessage {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 835 @column 0 */
  flushHeaders(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 908 @column 0 */
  setNoDelay(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 923 @column 0 */
  setSocketKeepAlive(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    method: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    path: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    protocol: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    readonly reusedSocket: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    readonly socket: Socket;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    host: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/http/messages.go @line 428 @column 0 */
    maxHeadersCount: number;
  }
  export { ClientRequest };
}

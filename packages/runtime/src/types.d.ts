declare module 'web-streams-polyfill/ponyfill/es2018' {
  import {
    AbortSignal,
    ByteLengthQueuingStrategy,
    CountQueuingStrategy,
    QueuingStrategy,
    QueuingStrategyInit,
    ReadableByteStreamController,
    ReadableStream as BaseReadableStream,
    ReadableStreamAsyncIterator,
    ReadableStreamBYOBReadResult,
    ReadableStreamBYOBReader,
    ReadableStreamBYOBRequest,
    ReadableStreamDefaultController,
    ReadableStreamDefaultReadResult,
    ReadableStreamDefaultReader,
    ReadableStreamIteratorOptions,
    ReadableWritablePair,
    StreamPipeOptions,
    TransformStream,
    TransformStreamDefaultController,
    Transformer,
    TransformerFlushCallback,
    TransformerStartCallback,
    TransformerTransformCallback,
    UnderlyingByteSource,
    UnderlyingByteSourcePullCallback,
    UnderlyingByteSourceStartCallback,
    UnderlyingSink,
    UnderlyingSinkAbortCallback,
    UnderlyingSinkCloseCallback,
    UnderlyingSinkStartCallback,
    UnderlyingSinkWriteCallback,
    UnderlyingSource,
    UnderlyingSourceCancelCallback,
    UnderlyingSourcePullCallback,
    UnderlyingSourceStartCallback,
    WritableStream,
    WritableStreamDefaultController,
    WritableStreamDefaultWriter,
  } from 'web-streams-polyfill';

  import { Readable } from 'node:stream';

  export {
    AbortSignal,
    ByteLengthQueuingStrategy,
    CountQueuingStrategy,
    QueuingStrategy,
    QueuingStrategyInit,
    ReadableByteStreamController,
    ReadableStream,
    ReadableStreamAsyncIterator,
    ReadableStreamBYOBReadResult,
    ReadableStreamBYOBReader,
    ReadableStreamBYOBRequest,
    ReadableStreamDefaultController,
    ReadableStreamDefaultReadResult,
    ReadableStreamDefaultReader,
    ReadableStreamIteratorOptions,
    ReadableWritablePair,
    StreamPipeOptions,
    TransformStream,
    TransformStreamDefaultController,
    Transformer,
    TransformerFlushCallback,
    TransformerStartCallback,
    TransformerTransformCallback,
    UnderlyingByteSource,
    UnderlyingByteSourcePullCallback,
    UnderlyingByteSourceStartCallback,
    UnderlyingSink,
    UnderlyingSinkAbortCallback,
    UnderlyingSinkCloseCallback,
    UnderlyingSinkStartCallback,
    UnderlyingSinkWriteCallback,
    UnderlyingSource,
    UnderlyingSourceCancelCallback,
    UnderlyingSourcePullCallback,
    UnderlyingSourceStartCallback,
    WritableStream,
    WritableStreamDefaultController,
    WritableStreamDefaultWriter,
  };
}

declare module 'text-decoding' {
  export { TextEncoder, TextDecoder } from 'node:util';
}

interface ArrayBuffer {
  /**
   * Read-only. The length of the ArrayBuffer (in bytes).
   */
  readonly byteLength: number;
  readonly resizable?: boolean;
  readonly maxByteLength?: number;

  /*
   * The ArrayBuffer constructor's length property whose value is 1.
   */
  length: number;
  /**
   * Returns a section of an ArrayBuffer.
   */
  slice(begin: number, end?: number): ArrayBuffer;
  resize?(newByteLength: number);
  readonly [Symbol.species]?: ArrayBuffer;
  readonly [Symbol.toStringTag]: 'ArrayBuffer';
}

interface ArrayBufferConstructor {
  readonly prototype: ArrayBuffer;
  new (
    byteLength: number,
    { maxByteLength }: { maxByteLength: number } = {}
  ): ArrayBuffer;
}

declare var SharedArrayBuffer: SharedArrayBufferConstructor;

interface SharedArrayBuffer {
  /**
   * Read-only. The length of the ArrayBuffer (in bytes).
   */
  readonly byteLength: number;
  readonly growable?: boolean;
  readonly maxByteLength?: number;

  /*
   * The SharedArrayBuffer constructor's length property whose value is 1.
   */
  length: number;
  /**
   * Returns a section of an SharedArrayBuffer.
   */
  slice(begin: number, end?: number): SharedArrayBuffer;
  grow?(newByteLength: number);
  readonly [Symbol.species]?: SharedArrayBuffer;
  readonly [Symbol.toStringTag]: 'SharedArrayBuffer';
}

interface SharedArrayBufferConstructor {
  readonly prototype: SharedArrayBuffer;
  new (
    byteLength: number,
    { maxByteLength }: { maxByteLength: number } = {}
  ): SharedArrayBuffer;
}

declare var SharedArrayBuffer: SharedArrayBufferConstructor;

declare module 'native:@grexie/workers/crypto' {
  declare class Hash {
    readonly algorithm: string;
    constructor(algorithm: string);
    update(buffer: ArrayBufferLike): void;
    digest(buffer: ArrayBufferLike): ArrayBufferLike;
  }
}

declare module 'native:@grexie/workers/stream' {
  declare interface ReadableStream {
    read(buffer: ArrayBufferLike): Promise<number>;
    close(): Promise<void>;
  }

  declare interface WritableStream {
    write(buffer: ArrayBufferLike): Promise<void>;
    close(): Promise<void>;
  }

  declare interface Stream {
    read(buffer: ArrayBufferLike): Promise<number>;
    write(buffer: ArrayBufferLike): Promise<void>;
    close(): Promise<void>;
  }

  export { ReadableStream, WritableStream, Stream };
}

declare module 'tty-browserify' {
  export * from 'node:tty';
}

declare module 'browserify-zlib' {
  export * from 'node:zlib';
}

declare module 'crypto-browserify' {
  export * from 'node:crypto';
}

declare module 'native:@grexie/workers/fs' {
  import { ReadableStream } from 'native:@grexie/workers/stream';

  export class Stats {
    isFile(): boolean;
    isDirectory(): boolean;
    readonly contentType: string;
    readonly mtimeMs: number;
    readonly size: number;
  }

  export class File {
    createReader(): ReadableStream;
    async readAll(): Promise<ArrayBufferLike>;
    readAllSync(): ArrayBufferLike;
    close(): void;
  }

  export class FS {
    open(file: string): File;

    async readFile(file: string): Promise<ArrayBufferLike>;
    readFileSync(file: string): ArrayBufferLike;
    async stat(file: string): Promise<Stats>;
    statSync(file: string): Stats;
  }

  interface LocalFSConstructor {
    prototype: FS;
    new (path: string): FS;
  }

  export const LocalFS: LocalFSConstructor;

  interface HttpFSConstructor {
    prototype: FS;
    new (url: string): FS;
  }

  export const HttpFS: HttpFSConstructor;
}

declare module 'task' {
  import { Readable, Writable } from 'stream';
  import { HttpServer } from 'native:@grexie/workers/http';
  import { FS } from 'native:@grexie/workers/fs';

  declare class Task {
    setImmediate(callback: () => void): NodeJS.Immediate;
    clearImmediate(immediate: NodeJS.Immediate): void;
    setTimeout(callback: () => void, duration: number): NodeJS.Timeout;
    clearTimeout(timeout: NodeJS.Timeout): void;

    readonly http: HttpServer;
    readonly fs: FS;
  }

  declare class Process {
    readonly stdin: Readable;
    readonly stdout: Writable;
    readonly stderr: Writable;

    readonly platform: string;

    readonly env: Record<string, string>;
    cwd(): string;
  }

  declare const task: Task;
  declare const process: Process;

  export { task, process, Task, Process };
}

declare module 'native:@grexie/workers/net' {
  import { Stream } from 'native:@grexie/workers/stream';

  declare interface Net {
    dial(network: string, address: string): Promise<Conn>;
    dialTimeout(
      network: string,
      address: string,
      timeout: number
    ): Promise<Conn>;
    listen(network: string, address: string): Promise<Listener>;
  }

  declare class HostNet implements Net {
    async dial(network: string, address: string): Promise<Conn>;
    async dialTimeout(
      network: string,
      address: string,
      timeout: number
    ): Promise<Conn>;
    async listen(network: string, address: string): Promise<Listener>;
  }

  declare interface Conn extends Stream {
    readonly localAddress: Address;
    readonly remoteAddress: Address;
    closeRead(): Promise<void>;
    closeWrite(): Promise<void>;
    setDeadline(deadline: Date): void;
    setReadDeadline(deadline: Date): void;
    setWriteDeadline(deadline: Date): void;
  }

  declare interface Listener {
    accept(): Promise<Conn>;
    close(): Promise<void>;
    readonly address: Address;
  }

  declare interface Address {
    readonly network: string;
    readonly address: string;
  }
}

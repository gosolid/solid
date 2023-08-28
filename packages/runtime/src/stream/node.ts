import {
  Stream as Stream,
  Readable as BaseReadable,
  Writable,
  ReadableOptions,
} from 'readable-stream';
import type { ReadableStream } from 'web-streams-polyfill/ponyfill/es2018';
import { ReadableWebToNodeStream } from 'readable-web-to-node-stream';

declare global {
  type Readable = BaseReadable & {
    fromWeb(stream: ReadableStream): Readable;
    new (options?: ReadableOptions): Readable;
  };

  var Readable: Readable;
}

export const Readable: Readable = BaseReadable as any;

Readable.fromWeb = (stream: ReadableStream) => {
  return new ReadableWebToNodeStream(stream) as unknown as Readable;
};

export { Stream, Writable, ReadableOptions };

export default Stream;

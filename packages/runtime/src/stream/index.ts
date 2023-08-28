import {
  AbortSignal,
  ByteLengthQueuingStrategy,
  CountQueuingStrategy,
  QueuingStrategy,
  QueuingStrategyInit,
  ReadableByteStreamController,
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
  WritableStreamDefaultController,
  WritableStreamDefaultWriter,
  ReadableStream,
  WritableStream,
} from 'web-streams-polyfill/ponyfill/es2018';

import { Readable } from 'stream';

declare global {
  interface ReadableStream {
    toNode(): typeof Readable;
  }
}

ReadableStream.prototype.toNode = function () {
  const stream = Readable.fromWeb(this);
  stream.setMaxListeners(20);
  return stream;
};

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

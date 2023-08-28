export * from 'crypto-browserify';

import { BinaryToTextEncoding } from 'crypto-browserify';
import { Hash as NativeHash } from 'native:@grexie/workers/crypto';

class Hash {
  readonly #hash: NativeHash;

  constructor(hash: NativeHash) {
    this.#hash = hash;
  }

  update(buffer: ArrayBufferLike | Buffer | string) {
    if (Buffer.isBuffer(buffer)) {
      buffer = buffer.buffer;
    }
    if (typeof buffer === 'string') {
      buffer = Buffer.from(buffer).buffer;
    }
    this.#hash.update(buffer);
    return this;
  }

  digest(): Buffer;
  digest(encoding?: BinaryToTextEncoding): string;
  digest(encoding?: BinaryToTextEncoding): Buffer | string {
    const digest = Buffer.from(this.#hash.digest(Buffer.from([])));
    if (encoding) {
      return digest.toString(encoding);
    } else {
      return digest;
    }
  }
}

export const createHash = (algorithm: string) => {
  const hash = new NativeHash(algorithm);
  return new Hash(hash);
};

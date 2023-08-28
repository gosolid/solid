import { ReadableStream } from '@grexie/workers/stream';
import { task } from 'task';
import path from 'path';
import {
  HttpFS,
  LocalFS,
  FS as NativeFS,
  File as NativeFile,
  Stats as NativeStats,
} from 'native:@grexie/workers/fs';

const MAX_BUFFER_SIZE = 8 * 1024 * 1024;

export class Stats {
  readonly #stats: NativeStats;

  constructor(stats: NativeStats) {
    if (!stats) {
      throw new Error('no stats');
    }
    this.#stats = stats;
  }

  isDirectory() {
    return this.#stats.isDirectory();
  }

  isFile() {
    return this.#stats.isFile();
  }

  get contentType() {
    return this.#stats.contentType;
  }

  get mtimeMs() {
    return this.#stats.mtimeMs;
  }

  get mtime() {
    return new Date(this.#stats.mtimeMs);
  }

  get size() {
    return this.#stats.size;
  }
}

export class File {
  readonly #file: NativeFile;

  constructor(file: NativeFile) {
    this.#file = file;
  }
}

export class Mount {
  readonly fs: FS;
  readonly path: string;

  constructor(fs: FS, path = '/') {
    this.fs = fs;
    this.path = path;
  }

  resolve(p: string): MountContext {
    p = path.join(this.path, p);
    return { fs: this.fs, path: p };
  }
}

export interface MountContext {
  readonly fs: FS;
  readonly path: string;
}

class FS {
  readonly #fs: NativeFS;
  readonly #mounts: Record<string, Mount> = {};

  constructor(fs: NativeFS) {
    this.#fs = fs;
  }

  mount(mountPoint: string, fs: FS, path: string = '/') {
    this.#mounts[mountPoint] = new Mount(fs, path);
  }

  unmount(mountPoint: string) {
    delete this.#mounts[mountPoint];
  }

  get mounts() {
    return Object.assign({}, this.#mounts);
  }

  resolveMount(path: string) {
    const candidates: string[] = [];
    for (const mount in this.#mounts) {
      if (path.startsWith(mount)) {
        candidates.push(mount);
      }
    }
    candidates.sort().reverse();
    const mount = this.#mounts[candidates[0]] ?? new Mount(this, '/');
    path = path.substring(candidates[0]?.length ?? 0);
    return mount.resolve(path);
  }

  open(path: string): File {
    const mount = this.resolveMount(path);
    if (mount.fs !== this) {
      return mount.fs.open(mount.path);
    }
    return new File(this.#fs.open(mount.path));
  }

  createReadableStream(path: string): ReadableStream<ArrayBufferLike> {
    const mount = this.resolveMount(path);
    if (mount.fs !== this) {
      return mount.fs.createReadableStream(mount.path);
    }

    const reader = this.#fs.open(mount.path).createReader();
    let buffer = new ArrayBuffer(10 * 1024, {
      maxByteLength: 10 * 1024,
    });
    let view = new Uint8Array(buffer);
    const stream = new ReadableStream<ArrayBufferLike>({
      async cancel() {
        await reader.close();
      },
      async pull(controller) {
        try {
          const n = await reader.read(buffer);
          controller.enqueue(view.slice(0, n));

          const nextByteLength = Math.min(
            buffer.byteLength * 16,
            buffer.maxByteLength ?? MAX_BUFFER_SIZE
          );

          if (n == buffer.byteLength && buffer.byteLength !== nextByteLength) {
            if (buffer.resizable) {
              buffer.resize!(nextByteLength);
            } else {
              buffer = new ArrayBuffer(nextByteLength);
              view = new Uint8Array(buffer);
            }
          }
        } catch (err: any) {
          console.info('read error', path, err);
          if (err.message === 'EOF') {
            await reader.close();
            controller.close();
          } else {
            controller.error(err);
          }
        }
      },
    });
    return stream;
  }

  createReadStream(path: string) {
    return this.createReadableStream(path).toNode();
  }

  async readFile(
    path: string,
    callback?: (err: any, buffer?: Buffer) => void
  ): Promise<Buffer | undefined> {
    console.info('readFile', path, this.#fs);
    try {
      const mount = this.resolveMount(path);
      if (mount.fs !== this) {
        return mount.fs.readFile(mount.path, callback);
      }
      const arrayBuffer = await this.#fs.readFile(mount.path);
      const buffer = Buffer.from(arrayBuffer);

      if (callback) {
        callback(null, buffer);
      } else {
        return buffer;
      }
    } catch (err) {
      if (callback) {
        callback(err);
      } else {
        throw err;
      }
    }
  }

  async stat(
    path: string,
    callback?: (err: any, stats?: Stats) => void
  ): Promise<Stats | undefined> {
    console.info('stat', path, this.#fs);
    try {
      const mount = this.resolveMount(path);
      if (mount.fs !== this) {
        const stats = await mount.fs.stat(mount.path, callback);
        return stats;
      }
      const s = await this.#fs.stat(mount.path);
      const stats = new Stats(s);
      if (callback) {
        callback(null, stats);
      } else {
        return stats;
      }
    } catch (err: any) {
      console.info('stat error', err);
      err = new Error('ENOENT: file not found');
      err.code = 'ENOENT';
      if (callback) {
        callback(err);
      } else {
        throw err;
      }
    }
  }

  get ReadStream() {
    const nothing = function () {};
    nothing.prototype = Object.create({});
    return nothing;
  }
}

export const createHttpFileSystem = (url: string) => new FS(new HttpFS(url));
export const createLocalFileSystem = (path: string) =>
  new FS(new LocalFS(path));

export default new FS(task.fs);

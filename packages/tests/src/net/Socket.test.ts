import { createServer, connect } from 'net';

describe('net.Socket', () => {
  it('should listen on socket', async () => {
    const server = createServer(socket => {
      console.info(socket);
    });

    await new Promise<void>((resolve, reject) => {
      server.once('error', reject);

      server.listen(0, 'localhost', () => {
        resolve();
      });
    });

    await new Promise<void>((resolve, reject) => {
      server.once('error', err => {
        reject(err);
      });

      server.once('close', () => {
        resolve();
      });

      server.close();
    });
  });

  it('should connect to listening socket', async () => {
    const server = createServer();

    await new Promise<void>((resolve, reject) => {
      server.once('error', reject);

      server.listen(3030, 'localhost', () => {
        resolve();
      });
    });

    const conn = connect((server.address() as any).port, 'localhost');

    await Promise.all([
      new Promise<void>((resolve, reject) => {
        server.on('connection', conn => {
          conn.on('data', chunk => {
            if (chunk.toString() === 'EHLO') {
              resolve();
            } else {
              reject(new Error('invalid response received'));
            }
          });
          conn.resume();
          conn.write('HELO');
          conn.end();
        });
      }),
      new Promise<void>((resolve, reject) => {
        conn.once('error', err => {
          reject(err);
        });

        conn.once('connect', () => {
          resolve();
        });
      }),
      new Promise<void>((resolve, reject) => {
        conn.once('error', err => reject(err));

        conn.once('data', chunk => {
          if (chunk.toString() === 'HELO') {
            conn.write('EHLO');
            conn.end();
            resolve();
          } else {
            reject(new Error('invalid request received'));
          }
        });

        conn.resume();
      }),
    ]);

    await new Promise<void>((resolve, reject) => {
      server.once('error', err => {
        reject(err);
      });

      server.once('close', () => {
        resolve();
      });

      server.close();
    });
  });
});

import { createServer, request } from 'http';

describe.skip('http.request', () => {
  it('should create a get request', async () => {
    const server = createServer((_, res) => {
      res.write('hello world');
      res.end();
    });

    await new Promise<void>(resolve =>
      server.listen(0, 'localhost', () => {
        resolve();
      })
    );

    await new Promise<void>((resolve, reject) => {
      const req = request(
        `http://localhost:${(server.address() as any).port}`,
        res => {
          const chunks: Buffer[] = [];
          res.on('error', err => {
            console.error(err);
          });
          res.on('data', chunk => {
            chunks.push(chunk);
          });
          res.on('end', () => {
            const data = Buffer.concat(chunks).toString();

            if (data === 'hello world') {
              resolve();
            } else {
              reject(new Error('invalid response: ' + data));
            }
          });
          res.resume();
        }
      );

      req.on('error', err => {
        console.error(err);
      });
      req.write('');
      req.end();
    });

    await new Promise<void>((resolve, reject) => {
      server.closeAllConnections();

      server.on('error', err => {
        console.error(err);
        reject(err);
      });

      server.on('close', () => {
        resolve();
      });

      server.close();
    });
  });

  it('should transport a large chunk of data', async () => {
    const bytes = 'ABCD'.repeat(10240);

    const server = createServer((req, res) => {
      let i = 0;
      const next = () => {
        res.write(bytes.slice(i, i + 1037), 'utf8', err => {
          i += 1037;

          if (i < bytes.length) {
            next();
          } else {
            process.nextTick(() => {
              res.end();
            });
          }
        });
      };
      next();
    });

    await new Promise<void>(resolve =>
      server.listen(0, '127.0.0.1', () => {
        resolve();
      })
    );

    await new Promise<void>((resolve, reject) => {
      const req = request(
        `http://localhost:${(server.address() as any).port}`,
        res => {
          const chunks: Buffer[] = [];
          res.on('data', chunk => {
            chunks.push(chunk);
          });
          res.on('end', () => {
            const data = Buffer.concat(chunks).toString();

            if (data === bytes) {
              resolve();
            } else {
              reject(new Error('invalid response: ' + data.toString()));
            }
          });
          res.resume();
        }
      );
      req.write('');
      req.end();
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
});

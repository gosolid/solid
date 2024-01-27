import { createServer } from 'http';
import { connect } from 'net';

describe('http.Server', () => {
    it('should start a new http server', async () => {
        const server = createServer();

        await new Promise<void>((resolve, reject) => {
            server.once('error', reject);

            server.listen(0, 'localhost', () => {
                resolve();
            })
        });

        await new Promise<void>((resolve, reject) => {
            server.once('error', (err) => {
                reject(err)
            });

            server.once('close', () => {
                resolve();
            });
            
            server.close();
        });
    })

    it('should serve a short response', async () => {
        const server = createServer((req, res) => {
            res.write('hello world');
            res.end();
        });

        server.on('error', (err) => console.error(err));

        await new Promise<void>((resolve, reject) => {
            server.once('error', reject);

            server.listen(0, 'localhost', () => {
                resolve();
            })
        });

        const conn = connect((server.address() as any).port, 'localhost');

        await Promise.all([
            new Promise<void>((resolve, reject) => {
                conn.once('error', (err) => {
                    reject(err);
                })

                conn.once('connect', () => {
                    conn.write('GET / HTTP/1.0\r\n\r\n');
                    conn.end();
                    resolve();
                })
            }),
            new Promise<void>((resolve, reject) => {
                conn.once('error', err => reject(err));

                const chunks: Buffer[] = [];

                conn.once('data', (chunk) => {
                    chunks.push(chunk);
                })

                conn.on('end', () => {
                    const data = Buffer.concat(chunks).toString().split(/\r\n/g);
                    if (data[0] === 'HTTP/1.1 200 OK' && data[data.length - 1] === 'hello world') {
                        resolve();
                    } else {
                        reject(new Error('invalid response received: ' + data.join('\r\n')));
                    }
                })

                conn.resume();
            }),
        ])

        await new Promise<void>((resolve, reject) => {
            server.once('error', (err) => {
                reject(err)
            });

            server.once('close', () => {
                resolve();
            });
            
            server.close();
        });
    })

    it('should serve a short response 2', async () => {
        const server = createServer((req, res) => {
            res.write('hello world');
            res.end();
        });

        server.on('error', (err) => console.error(err));

        await new Promise<void>((resolve, reject) => {
            server.once('error', reject);

            server.listen(0, 'localhost', () => {
                resolve();
            })
        });

        const conn = connect((server.address() as any).port, 'localhost');

        await Promise.all([
            new Promise<void>((resolve, reject) => {
                conn.once('error', (err) => {
                    reject(err);
                })

                conn.once('connect', () => {
                    conn.write('GET / HTTP/1.0\r\n\r\n');
                    conn.end();
                    resolve();
                })
            }),
            new Promise<void>((resolve, reject) => {
                conn.once('error', err => reject(err));

                const chunks: Buffer[] = [];

                conn.once('data', (chunk) => {
                    chunks.push(chunk);
                })

                conn.on('end', () => {
                    const data = Buffer.concat(chunks).toString().split(/\r\n/g);
                    if (data[0] === 'HTTP/1.1 200 OK' && data[data.length - 1] === 'hello world') {
                        resolve();
                    } else {
                        reject(new Error('invalid response received: ' + data.join('\r\n')));
                    }
                })

                conn.resume();
            }),
        ])

        await new Promise<void>((resolve, reject) => {
            server.once('error', (err) => {
                reject(err)
            });

            server.once('close', () => {
                resolve();
            });
            
            server.close();
        });
    })
});
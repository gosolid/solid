import { createServer, request } from 'http';
import http from 'http';

describe('http.request', () => {
    it('should create a get request', async () => {
        const server = createServer((req, res) => {
            res.write('hello world');
            res.end();
        })

        await new Promise<void>((resolve) => server.listen(3030, '127.0.0.1', () => resolve()));

        await new Promise<void>((resolve, reject) => {
            const req = request(`http://127.0.0.1:${(server.address() as any).port}`, (res) => {
                res.on('data', chunk => {
                    if (chunk.toString() === 'hello world') {
                        resolve();
                    } else {
                        reject(new Error('invalid response: ' + chunk.toString()));
                    }
                })
                res.resume();
            });
            req.socket!.on('connect', () => {
                req.write('');
                req.end();
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

    it.skip('should transport a large chunk of data', async () => {
        const bytes = 'ABCD'.repeat(1024);

        const server = createServer((req, res) => {
            let i = 0;
            const next = () => {
                console.info('calling res.write');
                res.write(bytes.slice(i, 1024), 'utf8', (err) => {
                    console.info('callback', err)
                });
                
                i += 1024;

                if (i < bytes.length) {
                    console.info('calling next')
                    next();
                } else {
                    console.info('calling process.nextTick')
                    process.nextTick(() => {
                        console.info('calling res.end')
                        res.end()
                    });
                }
            }
            next();
        })

        await new Promise<void>((resolve) => server.listen(3030, '127.0.0.1', () => resolve()));

        await new Promise<void>((resolve, reject) => {
            const req = request(`http://127.0.0.1:${(server.address() as any).port}`, (res) => {
                const chunks: string[] = [];
                res.on('data', chunk => {
                    chunks.push(chunk.toString());
                    console.info('received chunk', chunk.length)
                })
                res.on('end', () => {
                    if (chunks.join('') === bytes) {
                        resolve();
                    } else {
                        reject(new Error('invalid response: ' + chunks.join('').toString()));
                    }
                })
                res.resume();
            });
            req.socket!.on('connect', () => {
                req.write('');
                req.end();
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
})
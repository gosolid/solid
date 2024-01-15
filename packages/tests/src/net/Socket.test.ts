import { createServer, connect } from 'net';

describe('net.Socket', () => {
    it('should listen on socket', async () => {
        const server = createServer((socket) => {
            console.info(socket);
        })


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

    it('should connect to listening socket', async () => {
        const server = createServer((socket) => {
            console.info(socket);
        })


        await new Promise<void>((resolve, reject) => {
            server.once('error', reject);

            server.listen(3000, 'localhost', () => {
                resolve();
            })
        });

        server.on('connection', conn => {
            conn.on('data', (chunk) => console.info(chunk));
            conn.write('EHLO');
            conn.resume();
        })

        const conn = connect(3000, 'localhost', () => {
            console.info("HERE")
        });

        conn.on('data', (chunk) => {
            console.info(chunk)
        })

        await new Promise<void>((resolve, reject) => {
            conn.on('error', (err) => {
                reject(err);
            })

            conn.on('connect', () => {
                resolve();
            })
        })


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
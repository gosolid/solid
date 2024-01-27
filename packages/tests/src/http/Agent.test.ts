import { Agent } from 'http';
import { Socket, createServer } from 'net';
import { expect } from 'chai';

describe('http.Agent', () => {
    it('should create a new agent', () => {
        const agent = new Agent();
        expect(agent).to.be.instanceOf(Agent);
    });

    it('should return a canonical name', () => {
        const agent = new Agent();
        const name = agent.getName({
            host: 'localhost',
            port: '3000',
            localAddress: '127.0.0.1',
        });
        expect(name).to.equal('localhost:3000:127.0.0.1');
    })

    it('should return a canonical name for ipv4', () => {
        const agent = new Agent();
        const name = agent.getName({
            host: 'localhost',
            port: '3000',
            localAddress: '127.0.0.1',
            family: 4
        });
        expect(name).to.equal('localhost:3000:127.0.0.1:4');
    })

    it('should return a canonical name for ipv6', () => {
        const agent = new Agent();
        const name = agent.getName({
            host: 'localhost',
            port: '3000',
            localAddress: '::0',
            family: 6
        });
        expect(name).to.equal('localhost:3000:::0:6');
    })

    it('should create a connection', async () => {
        const server = createServer();

        await new Promise<void>((resolve, reject) => {
            server.once('error', reject);

            server.listen(0, 'localhost', () => {
                resolve();
            });
        });

        server.on('connection', (conn) => {
            conn.write('');
            conn.end();
        });

        const agent = new Agent();
        const conn = agent.createConnection({
            port: (server.address() as any).port,
            host: 'localhost',
            family: 4,
        });
        expect(conn).to.be.instanceOf(Socket);

        await new Promise<void>((resolve) => conn.on('close', () => {
            resolve();
        }))

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
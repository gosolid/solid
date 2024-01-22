import { expect } from 'chai';
import { createHmac } from 'crypto';

describe('crypto.hmac', () => {
    it('should create a sha256 hmac', () => {
        const hash = createHmac('sha256', '123456');
        hash.update('123456');
        expect(hash.digest('hex')).to.equal('b8ad08a3a547e35829b821b75370301dd8c4b06bdd7771f9b541a75914068718');
    })

    it('should create a sha512 hmac', () => {
        const hash = createHmac('sha512', '123456');
        hash.update('123456');
        expect(hash.digest('hex')).to.equal('4899f48b7873797086fc392ed8074b34306f79145cf0f9d1757e806da2d43f3876b3c762f38015f2d3593a595ae607a6e0aa103a2a5fe502cf95051c9cd62ee1');
    })

    it('should create a sha256 hmac from buffer', () => {
        const hash = createHmac('sha256', '123456');
        hash.update(Buffer.from('123456'));
        expect(hash.digest('hex')).to.equal('b8ad08a3a547e35829b821b75370301dd8c4b06bdd7771f9b541a75914068718');
    })

    it('should create a ripemd160 hmac', () => {
        const hash = createHmac('ripemd160', '123456');
        hash.update('123456');
        expect(hash.digest('hex')).to.equal('636f943fa94a080170267c7fcc5047b1dce17a89');
    })
})
import { expect } from 'chai';

describe('crypto/web.getRandomValues', () => {
    it('should get random values', () => {
        const arrayBufferView = new Uint8Array(128);
        expect(crypto.getRandomValues(arrayBufferView)).to.equal(arrayBufferView);
    })

    it('should generate a Buffer from ArrayBufferView', () => {
        const arrayBufferView = new Uint8Array(128);
        expect(Buffer.from(crypto.getRandomValues(arrayBufferView))).to.be.instanceOf(Buffer);
    })

    it('should generate a Buffer from ArrayBuffer', () => {
        const arrayBufferView = new Uint8Array(128);
        expect(Buffer.from(crypto.getRandomValues(arrayBufferView).buffer)).to.be.instanceOf(Buffer);
    })

    it('should serialize to a string and then back to a Buffer', () => {
        const arrayBufferView = new Uint8Array(128);
        expect(Buffer.from(crypto.getRandomValues(arrayBufferView)).length).to.equal(128);
        expect(Buffer.from(Buffer.from(crypto.getRandomValues(arrayBufferView).buffer).toString('utf8'))).to.be.instanceOf(Buffer);
    })
});
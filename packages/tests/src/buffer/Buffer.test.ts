import { expect } from 'chai';

describe('buffer.Buffer', () => {
    it('should create a buffer', () => {
        const buffer = Buffer.alloc(4);
        expect(buffer).to.be.instanceOf(Uint8Array);
        expect(buffer[0]).to.equal(0);
        expect(buffer[1]).to.equal(0);
        expect(buffer[2]).to.equal(0);
        expect(buffer[3]).to.equal(0);
        expect(buffer.length).to.equal(4);
    })

    it('should create a buffer from a string', () => {
        const buffer = Buffer.from('hello world');
        expect(buffer).to.be.instanceOf(Uint8Array);
        expect(buffer.length).to.equal(11);
        expect(buffer.toString()).equal('hello world');
    })

    it('should slice a buffer', () => {
        const buffer = Buffer.from('hello world');
        expect(buffer.slice(0, 5).toString()).to.equal('hello');
    })

    it('should slice a buffer without end', () => {
        const buffer = Buffer.from('hello world');
        expect(buffer.slice(6).toString()).to.equal('world');
    })

    it('should slice a buffer from end', () => {
        const buffer = Buffer.from('hello world');
        expect(buffer.slice(-5).toString()).to.equal('world');
    })

    it('should read a uint32BE', () => {
        const buffer = Buffer.from([0x00, 0x01, 0x02, 0x03]);
        expect(buffer.readUInt32BE(0)).to.equal(0x00010203);
    })

    it('should read a uint32BE from offset', () => {
        const buffer = Buffer.from([0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07]);
        expect(buffer.readUInt32BE(4)).to.equal(0x04050607);
    })

    it('should write a uint32BE', () => {
        const buffer = Buffer.from([0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07]);
        buffer.writeUInt32BE(0x08091011);
        expect(buffer.readUInt32BE(0)).to.equal(0x08091011);
        expect(buffer.readUInt32BE(4)).to.equal(0x04050607);
    })
})
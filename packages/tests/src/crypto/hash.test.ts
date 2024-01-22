import { expect } from 'chai';
import { createHash } from 'crypto';

describe('crypto.hash', () => {
    it('should create a sha256 hash', () => {
        const hash = createHash('sha256');
        hash.update('123456');
        expect(hash.digest('hex')).to.equal('8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92');
    })

    it('should create a sha512 hash', () => {
        const hash = createHash('sha512');
        hash.update('123456');
        expect(hash.digest('hex')).to.equal('ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413');
    })

    it('should create a sha256 hash from buffer', () => {
        const hash = createHash('sha256');
        hash.update(Buffer.from('123456'));
        expect(hash.digest('hex')).to.equal('8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92');
    })

    it('should create a ripemd160 hash', () => {
        const hash = createHash('ripemd160');
        hash.update('123456');
        expect(hash.digest('hex')).to.equal('d8913df37b24c97f28f840114d05bd110dbb2e44');
    })
})
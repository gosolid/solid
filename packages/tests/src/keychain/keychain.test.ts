import keychain from 'keychain';
import { expect } from 'chai';

describe.only('keychain', () => {
  it('should create a new credential', () => {
    // console.info(keychain.has('solid'));
    const credential = keychain.create('solid');
    credential.data = Buffer.from('test password');

    expect(credential.data.toString()).to.equal('test password');
  }).timeout(20_000);
});

import { UserSalt } from '../hooks/useApollo';
import _scrypt from 'scrypt-js/scrypt';

export const scrypt = async (salt: UserSalt, password: string) => {
  const passwordBuffer = new TextEncoder().encode(password);

  const derivedKey = await _scrypt.scrypt(
    passwordBuffer,
    new Uint8Array(Buffer.from(salt.salt, 'base64').buffer),
    1024,
    8,
    1,
    salt.keyLength
  );

  return `${[salt.algorithm, salt.keyLength, salt.salt].join(
    '$'
  )}\$${Buffer.from(derivedKey).toString('base64')}`;
};

import { format, formatWithOptions } from 'util';
import { ansi } from 'tty';
import { expect } from 'chai';

describe('util.format', () => {
  const literal = ansi.cyan;
  const key = ansi.black.bold;
  const other = ansi.white.dim;
  const syntax = ansi.white.dim.italic;
  const errorStack = ansi.white.dim;
  const errorMessage = ansi.red.bold;

  it('should return an empty string', () => {
    expect(format()).to.equal('');
  });

  it('should format first argument', () => {
    expect(format('%s:%s', 'foo')).to.equal('foo:%s');
  });

  it('should format remaining arguments', () => {
    expect(format('%s:%s', 'foo', 'bar', 'baz')).to.equal('foo:bar baz');
  });

  it('should format numbers', () => {
    expect(format(1, 2, 3)).to.equal('1 2 3');
  });

  it('should not format single argument', () => {
    expect(format('%% %s')).to.equal('%% %s');
  });

  it('should format an array', () => {
    expect(format([1, 2, 3])).to.equal('array [1, 2, 3]');
  });

  it('should format a large array', () => {
    expect(format(Array.from({ length: 100 }, (_, i) => `hello ${i}`))).to
      .equal(`array [
  'hello 0', 'hello 1', 'hello 2', 'hello 3', 'hello 4', 'hello 5', 'hello 6', 
  'hello 7', 'hello 8', 'hello 9', 'hello 10', 'hello 11', 'hello 12', 'hello 13', 
  'hello 14', 'hello 15', 'hello 16', 'hello 17', 'hello 18', 'hello 19', 
  'hello 20', 'hello 21', 'hello 22', 'hello 23', 'hello 24', 'hello 25', 
  'hello 26', 'hello 27', 'hello 28', 'hello 29', 'hello 30', 'hello 31', 
  'hello 32', 'hello 33', 'hello 34', 'hello 35', 'hello 36', 'hello 37', 
  'hello 38', 'hello 39', 'hello 40', 'hello 41', 'hello 42', 'hello 43', 
  'hello 44', 'hello 45', 'hello 46', 'hello 47', 'hello 48', 'hello 49', 
  'hello 50', 'hello 51', 'hello 52', 'hello 53', 'hello 54', 'hello 55', 
  'hello 56', 'hello 57', 'hello 58', 'hello 59', 'hello 60', 'hello 61', 
  'hello 62', 'hello 63', 'hello 64', 'hello 65', 'hello 66', 'hello 67', 
  'hello 68', 'hello 69', 'hello 70', 'hello 71', 'hello 72', 'hello 73', 
  'hello 74', 'hello 75', 'hello 76', 'hello 77', 'hello 78', 'hello 79', 
  'hello 80', 'hello 81', 'hello 82', 'hello 83', 'hello 84', 'hello 85', 
  'hello 86', 'hello 87', 'hello 88', 'hello 89', 'hello 90', 'hello 91', 
  'hello 92', 'hello 93', 'hello 94', 'hello 95', 'hello 96', 'hello 97', 
  'hello 98', 'hello 99'
]`);
  });

  it('should format an object', () => {
    class A {
      hello() {
        return 'there';
      }
    }

    class B extends A {
      hello() {
        return 'three';
      }
      get hello2() {
        return 'hello there';
      }
    }

    class C extends B {
      hello() {
        return 'world';
      }

      get hello2() {
        return 'hello world';
      }
    }

    const value = new C();

    (value as any).a = 1;
    (value as any).b = 2;
    Object.defineProperty(value, 'hello2', {
      get: () => 'hello three',
    });

    expect(formatWithOptions({ showHidden: true }, value)).to.equal(`object C {
  a: 1
  b: 2
  get hello2(): 'hello three' overrides C, B
  extends C
    function hello(...) overrides B, A
  extends B
  extends A
}`);
  });

  it('should format an object without extends', () => {
    const value = {};

    (value as any).a = 1;
    (value as any).b = 2;
    Object.defineProperty(value, 'hello', {
      get: () => 'hello world',
    });

    expect(formatWithOptions({ showHidden: true }, value)).to.equal(`object {
  a: 1
  b: 2
  get hello(): 'hello world'
}`);
  });

  it('should format an object showing only enumerable properties', () => {
    const value = {};

    (value as any).a = 1;
    (value as any).b = 2;
    Object.defineProperty(value, 'hello', {
      get: () => 'hello world',
    });

    expect(format(value)).to.equal(`object {
  a: 1
  b: 2
}`);
  });

  it('should format empty objects on one line', () => {
    const value = {};

    expect(format(value)).to.equal(`object {}`);
  });

  it('should format embedded objects', () => {
    const value = {
      test: {
        hello: 'world',
      },
    };

    expect(format(value)).to.equal(`object {
  test: object {
    hello: 'world'
  }
}`);
  });

  it('should format embedded objects in prototype chain', () => {
    class A {
      get test() {
        return { hello: 'world' };
      }
    }
    const value = new A();

    expect(formatWithOptions({ showHidden: true }, value)).to.equal(`object A {
  extends A
    get test(): object {
      hello: 'world'
    }
}`);
  });

  it('should format an object with colors', () => {
    class A {
      hello() {
        return 'there';
      }
    }

    class B extends A {
      hello() {
        return 'three';
      }
      get hello2() {
        return 'hello there';
      }
    }

    class C extends B {
      hello() {
        return 'world';
      }

      get hello2() {
        return 'hello world';
      }
    }

    const value = new C();

    (value as any).a = 1;
    (value as any).b = 2;
    Object.defineProperty(value, 'hello2', {
      get: () => 'hello three',
    });

    expect(formatWithOptions({ colors: true, showHidden: true }, value)).to
      .equal(`${syntax('object')} ${key('C')} ${other('{')}
  ${key('a')}${other(':')} ${literal('1')}
  ${key('b')}${other(':')} ${literal('2')}
  ${syntax('get')} ${key('hello2')}${other('()')}${other(':')} ${other(
      "'"
    )}${literal('hello three')}${other("'")} ${syntax('overrides')} ${other(
      'C'
    )}${other(', ')}${other('B')}
  ${syntax('extends')} ${other('C')}
    ${syntax('function')} ${key('hello')}${other('(...)')} ${syntax(
      'overrides'
    )} ${other('B')}${other(', ')}${other('A')}
  ${syntax('extends')} ${other('B')}
  ${syntax('extends')} ${other('A')}
${other('}')}`);
  });

  it('should format a date', () => {
    const date = new Date('2023-08-11T18:12:04.034+01:00');
    expect(format(date)).to.equal(`date ${date.toISOString()}`);
  });

  it('should format a date with color', () => {
    const date = new Date('2023-08-11T18:12:04.034+01:00');
    expect(formatWithOptions({ colors: true }, date)).to.equal(
      `${syntax('date')} ${literal(date.toISOString())}`
    );
  });

  it('should format an error', () => {
    const err = new TypeError('an error');

    expect(format(err)).to.equal(err.stack?.replace(/^\s+/gm, '  '));
  });

  it('should format an error with color', () => {
    const err = new TypeError('an error');

    const formatted = err.stack?.replace(/^\s+/gm, '  ');

    expect(formatWithOptions({ colors: true }, err)).to.equal(
      errorMessage(formatted?.split(/\n/g, 2)[0]) +
        '\n' +
        formatted?.split(/\n/g).map(errorStack).slice(1).join('\n')
    );
  });
});

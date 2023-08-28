import type { Writable } from 'stream';
import ansi from 'ansi-colors';
import strip from 'strip-ansi';

const TAB = '  ';
const BR = '\n';

function arrayAsString(
  arr: any[],
  depth = 0,
  seen: WeakSet<any>,
  maxDepth: number
): string {
  const _ = TAB.repeat(depth - 1);
  const __ = _ + TAB;

  let out: string[] = [];

  if (depth <= maxDepth) {
    const serialized = arr.map(
      value => `${serialize(value, depth, maxDepth, seen)}`
    );

    let line = '';
    while (serialized.length) {
      line += serialized.shift();
      if (serialized.length) {
        line += ', ';
      }
      if (strip(line).length >= 50) {
        out.push(line);
        line = '';
      }
    }
    if (line.length) {
      out.push(line);
    }
    if (out.length > 1) {
      out.unshift(__);
    }
  }

  return `${ansi.grey.italic('array')}${
    depth > maxDepth
      ? ''
      : ` ${ansi.grey('[')}${out.join(BR + __)}${
          (out.length > 1 ? BR + _ : '') + ansi.grey(']')
        }`
  }`;
}

function objectAsString(
  obj: any,
  depth = 0,
  seen: WeakSet<any>,
  maxDepth: number
): string {
  const _ = TAB.repeat(depth - 1);
  const __ = _ + TAB;

  let descriptors: Map<
    string,
    PropertyDescriptor & { constructors: Function[] }
  > = new Map();
  const prototypes: Function[] = [];

  let o = obj;
  do {
    if (o) {
      const ownDescriptors: Record<
        string | symbol | number,
        PropertyDescriptor
      > = Object.getOwnPropertyDescriptors(o);
      const constructor = o === obj ? null : o.constructor;
      for (const k in ownDescriptors) {
        if (descriptors.has(k)) {
          if (!descriptors.get(k)!.constructors.includes(constructor)) {
            descriptors.get(k)!.constructors.push(constructor);
          }
        } else {
          descriptors.set(
            k,
            Object.assign({}, ownDescriptors[k], {
              constructors: [constructor],
            })
          );
        }
      }
      if (constructor) {
        prototypes.push(constructor);
      }
    }
  } while ((o = Object.getPrototypeOf(o)));
  // delete descriptors['toString'];

  const entries = [...descriptors.entries()].map(([name, descriptor]) => {
    try {
      return [
        name.toString(),
        { constructors: descriptor.constructors, value: obj[name] },
      ];
    } catch (err) {
      return [name.toString(), err];
    }
  }) as [string, { constructors: Function[]; value: any }][];

  entries.sort((a, b) => {
    const order =
      prototypes.indexOf(a[1].constructors?.[0]) -
      prototypes.indexOf(b[1].constructors?.[0]);

    if (order === 0) {
      return entries.indexOf(a) - entries.indexOf(b);
    } else {
      return order;
    }
  });

  let previousConstructor: Function;
  return `${ansi.grey(
    `${ansi.grey.italic('object')} ${ansi.cyan(
      `${Object.getPrototypeOf(obj)?.constructor.name ?? ''}`
    )}${Object.getPrototypeOf(obj)?.constructor.name ? ' ' : ''}`
  )}${
    depth > maxDepth
      ? ''
      : `${ansi.grey('{')}${BR}${entries
          .map(([key, value]) => {
            if (key === 'constructor') {
              return [];
            }

            const descriptor = descriptors.get(key)!;
            let labels: string[] = [];
            if (
              !previousConstructor ||
              previousConstructor !== descriptor.constructors[0]
            ) {
              previousConstructor = descriptor.constructors[0];
              if (descriptor.constructors[0]) {
                for (const constructor of prototypes.splice(
                  0,
                  prototypes.indexOf(descriptor.constructors[0]) + 1
                )) {
                  labels.push(
                    `${ansi.grey.italic('extends')} ${ansi.grey(
                      constructor?.name
                    )}:${__}`
                  );
                }
              }
            }

            const overrides = descriptor.constructors.slice(
              descriptor.constructors.indexOf(previousConstructor) + 1
            );

            let accessors = [];

            if (descriptor.get) {
              accessors.push('get');
            } else if (descriptor.set) {
              accessors.push('set');
            }

            if (typeof value.value === 'function' && value.value.name === key) {
              return [
                ...labels.map(label => `${__}${label}`),
                `${__}${descriptor.constructors[0] ? __ : ''}${ansi.italic.grey(
                  'function'
                )} ${ansi.bold.black(key)}${ansi.grey('(...)')}${
                  overrides.length
                    ? ` ${ansi.grey.italic('overrides')} ${overrides
                        .map(x => ansi.grey(x.name))
                        .join(ansi.grey(', '))}`
                    : ''
                }`,
              ];
            }

            return [
              ...labels.map(label => `${__}${label}`),
              `${__}${descriptor.constructors[0] ? __ : ''}${ansi.italic.grey(
                accessors.join('/') + (accessors.length ? ' ' : '')
              )}${ansi.bold.black(key)}${ansi.grey(
                accessors.length ? '()' : ''
              )}: ${serialize(value.value, depth, maxDepth, seen)}${
                overrides.length
                  ? ` ${ansi.grey.italic('overrides')} ${overrides
                      .map(x => ansi.grey(x.name))
                      .join(ansi.grey(', '))}`
                  : ''
              }`,
            ];
          })
          .reduce((a, b) => [...a, ...b], [])
          .filter(x => !!x)
          .join(BR)}${BR}${_}${ansi.grey('}')}`
  }`;
}

function serialize(
  value: any,
  depth = 0,
  maxDepth = 2,
  seen = new WeakSet<any>()
) {
  const _ = TAB.repeat(depth);
  const __ = _ + TAB;

  if (value === null) {
    return `${ansi.cyan('null')}`;
  } else if (value instanceof Error) {
    return ansi.red(
      value.stack
        ? `${ansi.bold(value.stack?.split(/\n/g)[0])}\n${ansi.grey(
            value.stack
              ?.split(/\n/)
              .map(x => x.replace(/^\s+/, __))
              .slice(1)
              .join('\n')
          )}`
        : value.message
        ? `${ansi.bold(value.message)}`
        : `${value}`
    );
  } else if (value instanceof Date) {
    return ansi.grey(
      `${ansi.grey.italic('date')} ${ansi.cyan(value.toISOString())}`
    );
  } else if (typeof value === 'function') {
    return ansi.grey(
      `${ansi.grey.italic('function')} ${ansi.cyan(value.name)}${ansi.grey(
        '(...)'
      )}`
    );
  } else if (Array.isArray(value)) {
    if (seen.has(value)) {
      return ansi.grey('[Circular]');
    }
    seen.add(value);

    return arrayAsString(value, depth + 1, seen, maxDepth);
  } else if (typeof value === 'object') {
    if (seen.has(value)) {
      return ansi.grey('[Circular]');
    }
    seen.add(value);

    return objectAsString(value, depth + 1, seen, maxDepth);
  } else if (typeof value === 'number') {
    return ansi.cyan(`${value}`);
  } else if (typeof value === 'boolean') {
    return ansi.cyan(`${value}`);
  } else if (typeof value === 'string') {
    let out = value.replace(/\\/g, '\\\\').replace(/'/g, "\\'");
    return ansi.cyan(`${ansi.grey("'") + out + ansi.grey("'")}`);
  } else {
    return `${value}`;
  }
}

export class Console {
  readonly #stdout: Writable;
  readonly #stderr: Writable;

  constructor(stdout: Writable, stderr: Writable) {
    this.#stdout = stdout;
    this.#stderr = stderr;
  }

  #write(stream: Writable, msg: string) {
    stream.write(msg);
  }

  #serialize(msg: any, maxDepth: number): string {
    if (typeof msg === 'string') {
      return msg;
    } else {
      ansi.enabled = true;
      return `${serialize(msg, 0, maxDepth).replace(/\n\n/g, '')}`;
    }
  }

  #log(
    type: 'info' | 'dir' | 'error' | 'log' | 'debug' | 'warn' | 'trace',
    stream: Writable,
    ...msg: any[]
  ) {
    const out = [];
    for (const m of msg) {
      out.push(this.#serialize(m, type === 'dir' ? 3 : 1));
    }
    this.#write(stream, out.join(' ') + '\n');
  }

  info = (...msg: any[]) => {
    this.#log('info', this.#stdout, ...msg);
  };

  dir = (...msg: any[]) => {
    this.#log('dir', this.#stdout, ...msg);
  };

  error = (...msg: any[]) => {
    this.#log('error', this.#stderr, ...msg);
  };

  log = (...msg: any[]) => {
    this.#log('log', this.#stdout, ...msg);
  };

  debug = (...msg: any[]) => {
    this.#log('debug', this.#stdout, ...msg);
  };

  warn = (...msg: any[]) => {
    this.#log('warn', this.#stderr, ...msg);
  };

  trace = (...msg: any[]) => {
    this.#log('trace', this.#stderr, ...msg);
  };
}

export const console = new Console(process.stdout, process.stderr);
export default Console;

// console.info(module.exports);

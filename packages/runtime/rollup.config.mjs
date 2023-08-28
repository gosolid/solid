import nodeResolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import typescript from '@rollup/plugin-typescript';
import alias from '@rollup/plugin-alias';
import define from 'rollup-plugin-define';
import json from '@rollup/plugin-json';

const defaultExternals = {
  buffer: true,
  crypto: true,
  http: true,
  events: true,
  stream: true,
  net: true,
  path: true,
  fs: true,
  os: true,
  tty: true,
  zlib: true,
  task: true,
};

const config = (externals = {}) => {
  externals = Object.assign({}, defaultExternals, externals);

  return {
    external: id => {
      if (externals[id]) {
        return true;
      }

      if (/^(native:|@grexie\/workers($|\/))/.test(id) && !/\/node$/.test(id)) {
        console.info(id);
        return true;
      }

      return false;
    },
    output: {
      dir: './lib',
      entryFileNames: '[name].js',
      format: 'cjs',
      sourcemap: true,
      exports: 'named',
      //       footer: `
      // module.exports = exports.default ?? {};
      // Object.assign(module.exports, exports);
      // `,
    },
    plugins: [
      alias({
        entries: [
          externals.os && {
            find: '@grexie/workers/os/node',
            replacement: 'os',
          },
          externals.fs && {
            find: '@grexie/workers/fs/node',
            replacement: 'fs',
          },
          externals.path && {
            find: '@grexie/workers/path/node',
            replacement: 'path',
          },
        ].filter(x => !!x),
      }),
      json(),
      nodeResolve({ preferBuiltins: false }),
      commonjs(),
      typescript({ outputToFilesystem: true }),
    ],
  };
};

export default [
  {
    input: {
      console: './src/console.ts',
    },
    ...config(),
    plugins: [...config().plugins],
  },
  {
    input: {
      crypto: './src/crypto.ts',
    },
    ...config({ crypto: false }),
    plugins: [
      alias({
        entries: [{ find: 'crypto', replacement: 'crypto-browserify' }],
      }),
      ...config().plugins,
    ],
  },
  {
    input: {
      events: './src/events.ts',
    },
    ...config({ events: false }),
  },
  {
    input: {
      'fs/index': './src/fs/index.ts',
    },
    ...config(),
  },
  {
    input: {
      'fs/node': './src/fs/node.ts',
    },
    ...config(),
  },
  {
    input: {
      global: './src/global.ts',
    },
    ...config(),
  },
  {
    input: {
      os: './src/os.ts',
    },
    ...config(),
  },
  {
    input: {
      path: './src/path.ts',
    },
    ...config({ path: false }),
  },
  {
    input: {
      text: './src/text.ts',
    },
    ...config(),
  },
  {
    input: {
      timers: './src/timers.ts',
    },
    ...config(),
  },
  {
    input: {
      tty: './src/tty.ts',
    },
    ...config({ tty: false }),
  },
  {
    input: {
      zlib: './src/zlib.ts',
    },
    ...config({ zlib: false }),
  },
];

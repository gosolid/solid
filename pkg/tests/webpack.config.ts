import path from 'path';
import webpack, { Configuration } from 'webpack';
import TerserPlugin from 'terser-webpack-plugin';
import NodePolyfillPlugin from 'node-polyfill-webpack-plugin';
import { createRequire } from 'module';

const __filename = new URL(import.meta.url).pathname;
const __dirname = path.dirname(__filename);
const require = createRequire(__filename);

const config: Configuration = {
  context: path.resolve(__dirname, 'src'),
  mode: process.env.NODE_ENV === 'production' ? 'production' : 'development',
  devtool: 'inline-source-map',
  output: {
    path: path.resolve(__dirname, 'lib'),
    filename: '[name].js',
    chunkFormat: false,
  },
  target: 'node',
  externals: [
    {
      buffer: 'commonjs buffer',
      crypto: 'commonjs crypto',
      http: 'commonjs http',
      events: 'commonjs events',
      path: 'commonjs path',
      stream: 'commonjs stream',
      net: 'commonjs net',
      os: 'commonjs os',
      fs: 'commonjs fs',
      tty: 'commonjs tty',
      util: 'commonjs util',
      zlib: 'commonjs zlib',
      task: 'commonjs task',
      url: '{}',
      tls: '{}',
      https: '{}',
      fsevents: '{}',
    },
  ],
  externalsPresets: {
    node: false,
  },
  module: {
    rules: [
      {
        type: 'javascript/auto',
        test: /\.tsx?$/,
        exclude: [path.resolve(__dirname, 'webpack.config.ts'), /node_modules/],
        use: ['babel-loader'],
      },
    ],
  },
  resolve: {
    extensionAlias: {
      '.js': ['.ts', '.js'],
    },
    extensions: ['.ts', '.tsx', '.js', '.jsx', '.cjs', '.mjs'],
    fallback: {
      async_hooks: false,
      // crypto: false,
      'serve-static': false,
      // zlib: false,
      // path: require.resolve('path-browserify'),
      // tty: require.resolve('tty-browserify'),
      // debug: false,
      // net: 'empty',
    },
  },
  resolveLoader: {
    extensions: ['.cjs', '.js'],
  },
  plugins: [],
  stats: {
    modulesSpace: 0,
  },
  optimization: {
    usedExports: true,
    minimize: false,
    minimizer: [
      new TerserPlugin({
        extractComments: false,
        terserOptions: {
          output: {
            comments: false,
          },
        },
      }),
    ],
    concatenateModules: true,
    innerGraph: true,
    sideEffects: true,
    splitChunks: {
      chunks: 'all',
      minSize: 20000,
      minRemainingSize: 0,
      minChunks: 1,
      maxAsyncRequests: 30,
      maxInitialRequests: 30,
      enforceSizeThreshold: 50000,
      cacheGroups: {
        defaultVendors: false,
        default: false,
      },
    },
    runtimeChunk: false,
  },
};

const entries: Configuration[] = [
  {
    ...config,
    entry: { runner: ['./index'] },
  },
];

export default entries;

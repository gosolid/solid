export default {
  exclude: [/\.d\.tsx?$/],
  sourceMaps: true,
  test: [/\.[cm]?tsx?$/],
  presets: [
    '@babel/typescript',
    [
      '@babel/env',
      {
        targets: 'node 16',
        modules: false,
      },
    ],
  ],
};

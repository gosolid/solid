export default {
  exclude: [/\.d\.tsx?$/],
  sourceMaps: true,
  test: [/\.tsx?$/],
  presets: [
    '@babel/typescript',
    [
      '@babel/env',
      {
        targets: 'node 10',
        modules: false,
      },
    ],
  ],
};

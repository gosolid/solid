import { config } from 'dotenv-flow';
import path from 'path';
import { createRequire } from 'module';
import { withPages } from '@grexie/pages/next';

const __dirname = path.dirname(new URL(import.meta.url).pathname);

const require = createRequire(import.meta.url);

config({
  path: path.resolve(__dirname, '..', '..'),
  default_node_env: 'development',
  productionBrowserSourceMaps: true,
  output: 'export',
});

/** @type {import('next').NextConfig} */
let nextConfig = {
  trailingSlash: true,
  pageExtensions: ['ts', 'tsx', 'js', 'jsx'],
  reactStrictMode: true,
  serverRuntimeConfig: {
    domains: {
      web: process.env.DOMAIN,
      wildcard: process.env.WILDCARD_DOMAIN,
    },
  },
  publicRuntimeConfig: {
    api: {
      url: process.env.API_URL,
    },
  },
  webpack: config => {
    config.resolve.fallback = { fs: false };

    return config;
  },
};

nextConfig = await withPages(nextConfig, {
  pagesDir: 'pages',
  plugins: [
    '@grexie/pages-plugin-svg',
    //   [
    //     '@grexie/service-worker/plugin',
    //     {
    //       exclude: [],
    //       external: [
    //         new URL('https://fonts.googleapis.com'),
    //         new URL('https://fonts.gstatic.com'),
    //         new URL('https://app.posthog.com'),
    //         new URL('https://www.googletagmanager.com'),
    //       ],
    //     },
    //   ],
  ],
});

export default nextConfig;

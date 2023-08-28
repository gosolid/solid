// const Mocha = require('mocha');
const { task } = require("task");
// const { Inspector } = require('util');
// const express = require('express');
// const { createServer } = require('http');
// const { WebSocketServer } = require('ws');
// const buffer = require('buffer');

interface a {}

const main = async () => {
  console.clear();
  console.info("🧪 Grexie Workers Test Suite");

  // const inspector = new Inspector();

  // console.info(buffer);
  // console.info(buffer.Buffer.isBuffer(true));

  // const app = express();

  // app.disable('x-powered-by');
  // app.disable('etag');

  // app.get('/json/version', (req, res, next) => {
  //   res.json({
  //     Browser: 'solid/v1.0.0',
  //     'Protocol-Version': '1.1',
  //   });
  // });

  // app.get('/json', (req, res, next) => {
  //   res.json([
  //     {
  //       description: 'Grexie Solid',
  //       devtoolsFrontendUrl:
  //         'devtools://devtools/bundled/js_app.html?experiments=true&v8only=true&ws=127.0.0.1:9229/',
  //       devtoolsFrontendUrlCompat:
  //         'devtools://devtools/bundled/inspector.html?experiments=true&v8only=true&ws=127.0.0.1:9229/',
  //       faviconUrl: 'https://nodejs.org/static/images/favicons/favicon.ico',
  //       id: '1',
  //       title: 'Grexie Solid',
  //       type: 'node',
  //       url: 'file://',
  //       webSocketDebuggerUrl: 'ws://127.0.0.1:9229/',
  //     },
  //   ]);
  // });

  // const server = createServer(app);

  // server.listen(9229, '127.0.0.1');

  // const ws = new WebSocketServer({ server });

  // ws.on('error', err => {
  //   console.error(err);
  // });

  // ws.on('connection', ws => {
  //   console.info('connection', ws);

  //   ws.on('error', err => {
  //     console.error(err);
  //   });

  //   ws.on('message', message => {
  //     inspector.write(message);
  //   });

  //   inspector.on('data', data => {
  //     ws.send(data);
  //   });

  //   inspector.on('close', () => {
  //     ws.close();
  //   });

  //   ws.on('close', () => {
  //     inspector.destroy();
  //   });
  // });

  const Mocha = require("mocha");

  const mocha = new Mocha({
    ui: "bdd",
    color: true,
  });

  mocha.run();

  console.info("Hello World");

  setInterval(() => {
    console.log(task.stats);
  }, 2000);
};

main().catch((err) => console.error(err));

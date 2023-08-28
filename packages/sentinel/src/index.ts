import { HttpListener } from '@grexie/workers/http';
import express from 'express';
import fs from 'fs';
import events from 'events';
import path from 'path';
import { createServer } from 'http';

const app = express();
app.disable('etag');
// app.set('etag', 'weak');
// // app.set('static buffer size', Number.MAX_VALUE);

app.use((req, res, next) => {
  console.info(
    '📃',
    req.method,
    req.url,
    req.socket.remoteAddress + ':' + req.socket.remotePort
  );
  next();
});

// server.on('error', err => {
//   console.error(err);
// });

// server.on('request', req => {
//   req.on('error', err => console.error(err));
// });

const server = app.listen(3060, '127.0.0.1', () => {
  const { address, port } = server.address() as AddressInfo;
  console.info(`🚀 listening on ${address}:${port}`);
});

server.on('connection', socket => {
  console.info(
    '🕸️  connection to',
    socket.localAddress + ':' + socket.localPort,
    'from',
    socket.remoteAddress + ':' + socket.remotePort
  );
});

// if (process.env.NODE_ENV === 'development') {
app.use(async (req, res, next) => {
  let pathname = req.path;

  if (pathname.startsWith('/')) {
    pathname = pathname.substring(1);
  }

  try {
    let stats = await new Promise<fs.Stats>((resolve, reject) =>
      fs.stat(path.join('/dashboard', pathname), (err, stats) => {
        if (err) {
          reject(err);
          return;
        }

        resolve(stats);
      })
    );

    if (stats.isDirectory()) {
      if (pathname && !pathname.endsWith('/')) {
        return res.redirect(pathname + '/');
      }

      stats = await new Promise<fs.Stats>((resolve, reject) =>
        fs.stat(
          path.join('/dashboard', pathname, 'index.html'),
          (err, stats) => {
            if (err) {
              reject(err);
              return;
            }

            resolve(stats);
          }
        )
      );

      // console.info(stats);
      if (stats.isFile()) {
        pathname = path.join(req.path, 'index.html');
      } else {
        next();
        return;
      }
    }

    // const stream = fs.createReadStream(path.join('/dashboard', pathname));
    // stream.pipe(res.status(200).contentType(stats.contentType));

    res
      .status(200)
      .contentType(stats.contentType)
      .sendFile(path.join('/dashboard', pathname), { etag: false });
  } catch (err) {
    console.info(err);
    next(err);
  }
});
// } else {
//   app.use(
//     express.static('/dashboard', {
//       dotfiles: 'allow',
//     })
//   );
// }

app.use((req, res) => {
  res
    .status(404)
    .contentType('text/html')
    .sendFile('/dashboard/404/index.html');
});

app.use((err, req, res, next) => {
  res.status(500);
  res.set('Content-Type', 'text/plain');
  res.send(err.stack + '\n');
});

// app.use(async (req, res, next) => {
//   try {
//     const response = await fetch(`http://localhost:3070${req.url}`);
//     for (const k in response.headers.keys()) {
//       res.setHeader(k, response.headers.get(k)!);
//     }
//     res.status(response.status);
//     response.body.toNode().pipe(res);
//   } catch (err) {
//     next(err);
//   }
// });

// HttpListener.create('https://localhost', app);

// // WebSocketListener.create('https://localhost', conn => {
// //   conn.on('message', message => {});

// //   conn.on('close', () => {});
// // });

// // const conn = new PeerConnection('1', []);
// // console.info(conn);

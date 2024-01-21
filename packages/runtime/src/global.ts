import { task } from "task";
import console from "console";
import { Buffer } from "buffer";
import crypto from 'crypto/web';

Object.assign(global, {
  console,
  Buffer,
  setImmediate: task.setImmediate.bind(task),
  setTimeout: task.setTimeout.bind(task),
  clearTimeout: task.clearTimeout.bind(task),
  setInterval: task.setInterval.bind(task),
  clearInterval: task.clearInterval.bind(task),
  crypto: crypto,
  // clearImmediate: nativeTask.clearImmediate.bind(nativeTask),
});

// global.console.info(global);

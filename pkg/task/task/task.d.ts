/**
 * @moduleName task
@packageDescription
@module task*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 2 @column 0 */
declare module "task" {
  import type { Counter } from "util";
  import type { EventEmitter } from "events";
  import type { Writable, Readable } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 30 @column 0 */
  class Task {
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 152 @column 0 */
  wait(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 322 @column 0 */
  setTimeout(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 355 @column 0 */
  clearTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 363 @column 0 */
  setInterval(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 391 @column 0 */
  clearInterval(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 30 @column 0 */
    readonly timeCounter: Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 30 @column 0 */
    readonly stats: any;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 30 @column 0 */
    exitCode: number;
  }
  export { Task };
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
  class Process extends EventEmitter {
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 427 @column 0 */
  cwd(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 483 @column 0 */
  nextTick(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 504 @column 0 */
  exit(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly args: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly platform: string;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly task: Task;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly stdin: Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly version: string;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly env: Record<string, string>;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly versions: Task;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly stdout: Writable;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly stderr: Writable;
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 70 @column 0 */
    readonly argv: string[];
  }
  export { Process };
  /** @filename Users/tim/src/grexie/solid/pkg/task/task.go @line 532 @column 0 */
  class TaskStats {
  }
  export { TaskStats };
}

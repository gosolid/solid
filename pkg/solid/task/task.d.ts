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

  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 2 @column 0 */
declare module "task" {
  import type { Counter } from "util";
  import type { Writable, Readable } from "stream";

  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 31 @column 0 */
  class Task {
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 149 @column 0 */
  wait(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 313 @column 0 */
  setTimeout(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 346 @column 0 */
  clearTimeout(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 354 @column 0 */
  setInterval(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 382 @column 0 */
  clearInterval(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 31 @column 0 */
    exitCode: number;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 31 @column 0 */
    readonly timeCounter: Counter;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 31 @column 0 */
    readonly stats: any;
  }
  export { Task };
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
  class Process {
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 418 @column 0 */
  cwd(  ): string;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 474 @column 0 */
  nextTick(  ): any;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 490 @column 0 */
  exit(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly stdout: Writable;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly argv: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly stdin: Readable;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly env: Record<string, string>;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly args: string[];
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly task: Task;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly stderr: Writable;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly platform: string;
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 71 @column 0 */
    readonly versions: Task;
  }
  export { Process };
  /** @filename Users/tim/src/grexie/solid/pkg/solid/task.go @line 518 @column 0 */
  class TaskStats {
  }
  export { TaskStats };
}

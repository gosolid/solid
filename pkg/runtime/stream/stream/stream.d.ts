/**
 * @moduleName stream
@packageDescription
@module stream*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 3 @column 0 */
declare module "stream" {
  import type { EventEmitter } from "events";

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 57 @column 0 */
  function pipeline(  ): any;
  export { pipeline };
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 65 @column 0 */
  class Stream extends EventEmitter {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 160 @column 0 */
  destroy(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 65 @column 0 */
    readonly destroyed: boolean;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 65 @column 0 */
    readonly errored: any;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go @line 65 @column 0 */
    readonly closed: boolean;
  }
  export { Stream };
  export default Stream;
}

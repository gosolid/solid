/**
 * @moduleName events
@packageDescription
@module events*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 2 @column 0 */
declare module "events" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 47 @column 0 */
  class EventEmitter {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 62 @column 0 */
  on(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 62 @column 0 */
  on(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 84 @column 0 */
  once(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 106 @column 0 */
  emit(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 156 @column 0 */
  eventNames(  ): string[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 174 @column 0 */
  getMaxListeners(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 179 @column 0 */
  setMaxListeners(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 184 @column 0 */
  listenerCount(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 211 @column 0 */
  rawListeners(  ): any[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 211 @column 0 */
  rawListeners(  ): any[];
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 233 @column 0 */
  off(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 233 @column 0 */
  off(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 266 @column 0 */
  prependListener(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 288 @column 0 */
  prependListenerOnce(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/events/emitter.go @line 310 @column 0 */
  removeAllListeners(  ): void;
  }
  export default EventEmitter;
  export { EventEmitter };
}

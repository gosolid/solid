/**
 * @moduleName webrtc
@packageDescription
@module webrtc*/

/*
 * Solide JavaScript Engine
 * 
 * Copyright (C) 2010 - 2023 Grexie
 * All Rights Reserved
 */

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 2 @column 0 */
declare module "webrtc" {

  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 28 @column 0 */
  class WebRTC {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 99 @column 0 */
  createPeerConnection(  ): PeerConnection;
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 69 @column 0 */
  class SessionDescription {
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 76 @column 0 */
  class ICECandidate {
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 94 @column 0 */
  class PeerConnection {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 194 @column 0 */
  wait(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 251 @column 0 */
  createDataChannel(  ): DataChannel;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 270 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 309 @column 0 */
  createOffer(  ): SessionDescription;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 320 @column 0 */
  createAnswer(  ): SessionDescription;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 360 @column 0 */
  addICECandidate(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 94 @column 0 */
    readonly state: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 94 @column 0 */
    readonly channels: DataChannel[];
  }
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 386 @column 0 */
  class DataChannel {
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 412 @column 0 */
  close(  ): void;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 425 @column 0 */
  read(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 438 @column 0 */
  write(  ): number;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 386 @column 0 */
    readonly label: string;
  /** @filename Users/tim/src/grexie/solid/pkg/runtime/webrtc/webrtc.go @line 386 @column 0 */
    readonly isOpen: boolean;
  }
}

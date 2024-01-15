//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOStreamBufferQueueEntry
*/

type IOStreamBufferQueueEntry struct {
  BufferID uint `v8:"bufferID"`
  DataOffset uint `v8:"dataOffset"`
  DataLength uint `v8:"dataLength"`
  ControlOffset uint `v8:"controlOffset"`
  ControlLength uint `v8:"controlLength"`
  Reserved [3]uint `v8:"reserved"`
}

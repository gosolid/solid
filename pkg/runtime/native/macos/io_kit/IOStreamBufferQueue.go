//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOStreamBufferQueue
*/

type IOStreamBufferQueue struct {
  EntryCount uint `v8:"entryCount"`
  HeadIndex uint `v8:"headIndex"`
  TailIndex uint `v8:"tailIndex"`
  Reserved uint `v8:"reserved"`
  Queue [0]IOStreamBufferQueueEntry `v8:"queue"`
}

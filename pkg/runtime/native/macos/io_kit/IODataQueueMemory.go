//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IODataQueueMemory
*/

type IODataQueueMemory struct {
  QueueSize uint `v8:"queueSize"`
  Head uint `v8:"head"`
  Tail uint `v8:"tail"`
  Queue [1]IODataQueueEntry `v8:"queue"`
}

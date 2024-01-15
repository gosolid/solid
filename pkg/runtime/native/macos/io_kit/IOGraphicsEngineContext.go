//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOGraphicsEngineContext
*/

type IOGraphicsEngineContext struct {
  ContextLock int `v8:"contextLock"`
  State int `v8:"state"`
  Owner void `v8:"owner"`
  Version int `v8:"version"`
  StructSize int `v8:"structSize"`
  Reserved [8]int `v8:"reserved"`
}

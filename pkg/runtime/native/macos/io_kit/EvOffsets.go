//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.EvOffsets
*/

type EvOffsets struct {
  EvGlobalsOffset int `v8:"evGlobalsOffset"`
  EvShmemOffset int `v8:"evShmemOffset"`
}

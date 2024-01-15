//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.CDPMADescriptor
*/

type CDPMADescriptor struct {
  Reserved byte `v8:"reserved"`
  Control byte `v8:"control"`
  Adr byte `v8:"adr"`
  Tno byte `v8:"tno"`
  Point byte `v8:"point"`
  Address CDMSF `v8:"address"`
  Zero byte `v8:"zero"`
  P CDMSF `v8:"p"`
}

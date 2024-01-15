//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXEQElement
*/

type NXEQElement struct {
  Next int `v8:"next"`
  Sema int `v8:"sema"`
  Event NXEventPtr `v8:"event"`
}

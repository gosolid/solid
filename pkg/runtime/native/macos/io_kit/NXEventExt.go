//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXEventExt
*/

type NXEventExt struct {
  Payload NXEventPtr `v8:"payload"`
  Extension NXEventExtension `v8:"extension"`
}

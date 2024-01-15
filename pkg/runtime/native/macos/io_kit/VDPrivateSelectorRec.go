//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDPrivateSelectorRec
*/

type VDPrivateSelectorRec struct {
  Reserved uint `v8:"reserved"`
  Data [1]VDPrivateSelectorDataRec `v8:"data"`
}

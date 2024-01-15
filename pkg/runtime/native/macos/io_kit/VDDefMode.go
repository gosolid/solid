//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDDefMode
*/

type VDDefMode struct {
  CsID byte `v8:"csID"`
  Filler byte `v8:"filler"`
}

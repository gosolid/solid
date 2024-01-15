//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDFlagRecord
*/

type VDFlagRecord struct {
  CsMode byte `v8:"csMode"`
  Filler byte `v8:"filler"`
}

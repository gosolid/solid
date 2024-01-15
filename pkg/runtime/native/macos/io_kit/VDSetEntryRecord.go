//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDSetEntryRecord
*/

type VDSetEntryRecord struct {
  CsTable ColorSpec `v8:"csTable"`
  CsStart int16 `v8:"csStart"`
  CsCount int16 `v8:"csCount"`
}

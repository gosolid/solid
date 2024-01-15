//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDEntryRecord
*/

type VDEntryRecord struct {
  CsTable *byte `v8:"csTable"`
}

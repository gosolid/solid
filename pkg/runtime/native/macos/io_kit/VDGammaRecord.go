//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDGammaRecord
*/

type VDGammaRecord struct {
  CsGTable *byte `v8:"csGTable"`
}

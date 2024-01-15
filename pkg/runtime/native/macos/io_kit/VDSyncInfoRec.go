//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDSyncInfoRec
*/

type VDSyncInfoRec struct {
  CsMode byte `v8:"csMode"`
  CsFlags byte `v8:"csFlags"`
}

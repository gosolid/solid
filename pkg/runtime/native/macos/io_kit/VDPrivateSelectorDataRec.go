//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.VDPrivateSelectorDataRec
*/

type VDPrivateSelectorDataRec struct {
  PrivateParameters *void `v8:"privateParameters"`
  PrivateParametersSize uint64 `v8:"privateParametersSize"`
  PrivateResults *void `v8:"privateResults"`
  PrivateResultsSize uint64 `v8:"privateResultsSize"`
}

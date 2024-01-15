//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOPowerStateChangeNotification
*/

type IOPowerStateChangeNotification struct {
  PowerRef void `v8:"powerRef"`
  ReturnValue uint64 `v8:"returnValue"`
  StateNumber uint64 `v8:"stateNumber"`
  StateFlags uint64 `v8:"stateFlags"`
}

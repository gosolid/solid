//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOPMCalendarStruct
*/

type IOPMCalendarStruct struct {
  Year uint `v8:"year"`
  Month byte `v8:"month"`
  Day byte `v8:"day"`
  Hour byte `v8:"hour"`
  Minute byte `v8:"minute"`
  Second byte `v8:"second"`
  Selector byte `v8:"selector"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.IOTimingInformation
*/

type IOTimingInformation struct {
  AppleTimingID uint `v8:"appleTimingID"`
  Flags uint `v8:"flags"`
  DetailedInfo void `v8:"detailedInfo"`
}

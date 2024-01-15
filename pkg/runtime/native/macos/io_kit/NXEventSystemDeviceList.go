//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXEventSystemDeviceList
*/

type NXEventSystemDeviceList struct {
  Dev [16]NXEventSystemDevice `v8:"dev"`
}

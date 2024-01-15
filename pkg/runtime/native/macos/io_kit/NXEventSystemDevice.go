//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.NXEventSystemDevice
*/

type NXEventSystemDevice struct {
  Interface int `v8:"interface"`
  InterfaceAddr int `v8:"interfaceAddr"`
  DevType int `v8:"devType"`
  Id int `v8:"id"`
}

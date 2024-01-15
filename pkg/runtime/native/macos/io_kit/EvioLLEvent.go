//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.evioLLEvent
*/

type EvioLLEvent struct {
  SetCursor int `v8:"setCursor"`
  Type int `v8:"type"`
  Location IOGPoint `v8:"location"`
  Data void `v8:"data"`
  SetFlags int `v8:"setFlags"`
  Flags int `v8:"flags"`
}

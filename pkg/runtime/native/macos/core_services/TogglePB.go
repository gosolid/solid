//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TogglePB
*/

type TogglePB struct {
  TogFlags int64 `v8:"togFlags"`
  AmChars uint `v8:"amChars"`
  PmChars uint `v8:"pmChars"`
  Reserved [4]int64 `v8:"reserved"`
}

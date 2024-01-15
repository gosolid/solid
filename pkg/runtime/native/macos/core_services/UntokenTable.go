//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UntokenTable
*/

type UntokenTable struct {
  Len int16 `v8:"len"`
  LastToken int16 `v8:"lastToken"`
  Index [256]int16 `v8:"index"`
}

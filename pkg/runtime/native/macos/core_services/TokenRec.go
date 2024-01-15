//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TokenRec
*/

type TokenRec struct {
  TheToken int16 `v8:"theToken"`
  Position *byte `v8:"position"`
  Length int64 `v8:"length"`
  StringPosition *byte `v8:"stringPosition"`
}

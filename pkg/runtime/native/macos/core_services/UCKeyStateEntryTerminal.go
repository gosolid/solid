//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.UCKeyStateEntryTerminal
*/

type UCKeyStateEntryTerminal struct {
  CurState uint16 `v8:"curState"`
  CharData uint16 `v8:"charData"`
}

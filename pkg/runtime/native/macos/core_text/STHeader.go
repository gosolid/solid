//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.STHeader
*/

type STHeader struct {
  Filler byte `v8:"filler"`
  NClasses byte `v8:"nClasses"`
  ClassTableOffset uint16 `v8:"classTableOffset"`
  StateArrayOffset uint16 `v8:"stateArrayOffset"`
  EntryTableOffset uint16 `v8:"entryTableOffset"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.STXHeader
*/

type STXHeader struct {
  NClasses uint `v8:"nClasses"`
  ClassTableOffset uint `v8:"classTableOffset"`
  StateArrayOffset uint `v8:"stateArrayOffset"`
  EntryTableOffset uint `v8:"entryTableOffset"`
}

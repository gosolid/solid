//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.STXEntryTwo
*/

type STXEntryTwo struct {
  NewState uint16 `v8:"newState"`
  Flags uint16 `v8:"flags"`
  Index1 uint16 `v8:"index1"`
  Index2 uint16 `v8:"index2"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.STEntryTwo
*/

type STEntryTwo struct {
  NewState uint16 `v8:"newState"`
  Flags uint16 `v8:"flags"`
  Offset1 uint16 `v8:"offset1"`
  Offset2 uint16 `v8:"offset2"`
}

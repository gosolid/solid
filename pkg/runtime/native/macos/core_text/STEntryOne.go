//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.STEntryOne
*/

type STEntryOne struct {
  NewState uint16 `v8:"newState"`
  Flags uint16 `v8:"flags"`
  Offset1 uint16 `v8:"offset1"`
}

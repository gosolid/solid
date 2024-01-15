//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortSubtable
*/

type MortSubtable struct {
  Length uint16 `v8:"length"`
  Coverage uint16 `v8:"coverage"`
  Flags uint `v8:"flags"`
  U void `v8:"u"`
}

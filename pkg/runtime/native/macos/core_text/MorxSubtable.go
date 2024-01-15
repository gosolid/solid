//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxSubtable
*/

type MorxSubtable struct {
  Length uint `v8:"length"`
  Coverage uint `v8:"coverage"`
  Flags uint `v8:"flags"`
  U void `v8:"u"`
}

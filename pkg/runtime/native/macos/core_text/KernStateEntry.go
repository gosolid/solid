//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernStateEntry
*/

type KernStateEntry struct {
  NewState uint16 `v8:"newState"`
  Flags uint16 `v8:"flags"`
}

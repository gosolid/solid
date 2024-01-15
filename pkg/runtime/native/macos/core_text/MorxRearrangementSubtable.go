//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxRearrangementSubtable
*/

type MorxRearrangementSubtable struct {
  Header STXHeader `v8:"header"`
}

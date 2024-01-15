//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxLigatureSubtable
*/

type MorxLigatureSubtable struct {
  Header STXHeader `v8:"header"`
  LigatureActionTableOffset uint `v8:"ligatureActionTableOffset"`
  ComponentTableOffset uint `v8:"componentTableOffset"`
  LigatureTableOffset uint `v8:"ligatureTableOffset"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortLigatureSubtable
*/

type MortLigatureSubtable struct {
  Header STHeader `v8:"header"`
  LigatureActionTableOffset uint16 `v8:"ligatureActionTableOffset"`
  ComponentTableOffset uint16 `v8:"componentTableOffset"`
  LigatureTableOffset uint16 `v8:"ligatureTableOffset"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxContextualSubtable
*/

type MorxContextualSubtable struct {
  Header STXHeader `v8:"header"`
  SubstitutionTableOffset uint `v8:"substitutionTableOffset"`
}

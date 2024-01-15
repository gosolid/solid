//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortContextualSubtable
*/

type MortContextualSubtable struct {
  Header STHeader `v8:"header"`
  SubstitutionTableOffset uint16 `v8:"substitutionTableOffset"`
}

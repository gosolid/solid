//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortInsertionSubtable
*/

type MortInsertionSubtable struct {
  Header STHeader `v8:"header"`
}

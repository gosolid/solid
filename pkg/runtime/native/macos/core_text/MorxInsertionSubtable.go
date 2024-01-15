//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxInsertionSubtable
*/

type MorxInsertionSubtable struct {
  Header STXHeader `v8:"header"`
  InsertionGlyphTableOffset uint `v8:"insertionGlyphTableOffset"`
}

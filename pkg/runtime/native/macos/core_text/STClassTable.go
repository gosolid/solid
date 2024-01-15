//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.STClassTable
*/

type STClassTable struct {
  FirstGlyph uint16 `v8:"firstGlyph"`
  NGlyphs uint16 `v8:"nGlyphs"`
  Classes [1]byte `v8:"classes"`
}

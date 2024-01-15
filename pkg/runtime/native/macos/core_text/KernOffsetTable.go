//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernOffsetTable
*/

type KernOffsetTable struct {
  FirstGlyph uint16 `v8:"firstGlyph"`
  NGlyphs uint16 `v8:"nGlyphs"`
  OffsetTable [1]uint16 `v8:"offsetTable"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupSegment
*/

type SFNTLookupSegment struct {
  LastGlyph uint16 `v8:"lastGlyph"`
  FirstGlyph uint16 `v8:"firstGlyph"`
  Value [1]uint16 `v8:"value"`
}

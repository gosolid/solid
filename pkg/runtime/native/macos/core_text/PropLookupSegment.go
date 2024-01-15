//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.PropLookupSegment
*/

type PropLookupSegment struct {
  LastGlyph uint16 `v8:"lastGlyph"`
  FirstGlyph uint16 `v8:"firstGlyph"`
  Value uint16 `v8:"value"`
}

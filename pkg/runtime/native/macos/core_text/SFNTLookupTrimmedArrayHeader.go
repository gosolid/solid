//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupTrimmedArrayHeader
*/

type SFNTLookupTrimmedArrayHeader struct {
  FirstGlyph uint16 `v8:"firstGlyph"`
  Count uint16 `v8:"count"`
  ValueArray [1]uint16 `v8:"valueArray"`
}

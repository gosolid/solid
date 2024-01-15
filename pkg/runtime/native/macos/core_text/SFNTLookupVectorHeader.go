//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupVectorHeader
*/

type SFNTLookupVectorHeader struct {
  ValueSize uint16 `v8:"valueSize"`
  FirstGlyph uint16 `v8:"firstGlyph"`
  Count uint16 `v8:"count"`
  Values [1]byte `v8:"values"`
}

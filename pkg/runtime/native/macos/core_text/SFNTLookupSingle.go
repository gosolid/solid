//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupSingle
*/

type SFNTLookupSingle struct {
  Glyph uint16 `v8:"glyph"`
  Value [1]uint16 `v8:"value"`
}

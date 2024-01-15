//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.ALMXHeader
*/

type ALMXHeader struct {
  Version int `v8:"version"`
  Flags uint16 `v8:"flags"`
  NMasters uint16 `v8:"nMasters"`
  FirstGlyph uint16 `v8:"firstGlyph"`
  LastGlyph uint16 `v8:"lastGlyph"`
  Lookup SFNTLookupTable `v8:"lookup"`
}

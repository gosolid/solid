//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.BslnFormat3Part
*/

type BslnFormat3Part struct {
  StdGlyph uint16 `v8:"stdGlyph"`
  CtlPoints [32]int16 `v8:"ctlPoints"`
  MappingData SFNTLookupTable `v8:"mappingData"`
}

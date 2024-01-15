//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernIndexArrayHeader
*/

type KernIndexArrayHeader struct {
  GlyphCount uint16 `v8:"glyphCount"`
  KernValueCount byte `v8:"kernValueCount"`
  LeftClassCount byte `v8:"leftClassCount"`
  RightClassCount byte `v8:"rightClassCount"`
  Flags byte `v8:"flags"`
  KernValue [1]int16 `v8:"kernValue"`
  LeftClass [1]byte `v8:"leftClass"`
  RightClass [1]byte `v8:"rightClass"`
  KernIndex [1]byte `v8:"kernIndex"`
}

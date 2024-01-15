//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxSimpleArrayHeader
*/

type KerxSimpleArrayHeader struct {
  RowWidth uint `v8:"rowWidth"`
  LeftOffsetTable uint `v8:"leftOffsetTable"`
  RightOffsetTable uint `v8:"rightOffsetTable"`
  TheArray uint `v8:"theArray"`
  FirstTable [1]uint `v8:"firstTable"`
}

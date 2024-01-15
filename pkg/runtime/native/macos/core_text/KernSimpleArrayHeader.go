//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernSimpleArrayHeader
*/

type KernSimpleArrayHeader struct {
  RowWidth uint16 `v8:"rowWidth"`
  LeftOffsetTable uint16 `v8:"leftOffsetTable"`
  RightOffsetTable uint16 `v8:"rightOffsetTable"`
  TheArray uint16 `v8:"theArray"`
  FirstTable [1]uint16 `v8:"firstTable"`
}

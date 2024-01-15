//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.TrakTableData
*/

type TrakTableData struct {
  NTracks uint16 `v8:"nTracks"`
  NSizes uint16 `v8:"nSizes"`
  SizeTableOffset uint `v8:"sizeTableOffset"`
  TrakTable [1]TrakTableEntry `v8:"trakTable"`
}

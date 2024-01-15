//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxIndexArrayHeader
*/

type KerxIndexArrayHeader struct {
  Flags uint `v8:"flags"`
  RowCount uint16 `v8:"rowCount"`
  ColumnCount uint16 `v8:"columnCount"`
  RowIndexTableOffset uint `v8:"rowIndexTableOffset"`
  ColumnIndexTableOffset uint `v8:"columnIndexTableOffset"`
  KerningArrayOffset uint `v8:"kerningArrayOffset"`
  KerningVectorOffset uint `v8:"kerningVectorOffset"`
}

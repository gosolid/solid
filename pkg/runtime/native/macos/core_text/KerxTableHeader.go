//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxTableHeader
*/

type KerxTableHeader struct {
  Version int `v8:"version"`
  NTables uint `v8:"nTables"`
  FirstSubtable [1]uint `v8:"firstSubtable"`
}

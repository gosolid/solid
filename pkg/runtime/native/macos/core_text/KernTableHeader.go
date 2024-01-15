//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernTableHeader
*/

type KernTableHeader struct {
  Version int `v8:"version"`
  NTables int `v8:"nTables"`
  FirstSubtable [1]uint16 `v8:"firstSubtable"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernVersion0Header
*/

type KernVersion0Header struct {
  Version uint16 `v8:"version"`
  NTables uint16 `v8:"nTables"`
  FirstSubtable [1]uint16 `v8:"firstSubtable"`
}

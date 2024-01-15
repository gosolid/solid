//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.OpbdTable
*/

type OpbdTable struct {
  Version int `v8:"version"`
  Format uint16 `v8:"format"`
  LookupTable SFNTLookupTable `v8:"lookupTable"`
}

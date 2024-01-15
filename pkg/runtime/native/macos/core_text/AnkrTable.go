//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.AnkrTable
*/

type AnkrTable struct {
  Version uint16 `v8:"version"`
  Flags uint16 `v8:"flags"`
  LookupTableOffset uint `v8:"lookupTableOffset"`
  AnchorPointTableOffset uint `v8:"anchorPointTableOffset"`
}

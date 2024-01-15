//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupBinarySearchHeader
*/

type SFNTLookupBinarySearchHeader struct {
  UnitSize uint16 `v8:"unitSize"`
  NUnits uint16 `v8:"nUnits"`
  SearchRange uint16 `v8:"searchRange"`
  EntrySelector uint16 `v8:"entrySelector"`
  RangeShift uint16 `v8:"rangeShift"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernOrderedListHeader
*/

type KernOrderedListHeader struct {
  NPairs uint16 `v8:"nPairs"`
  SearchRange uint16 `v8:"searchRange"`
  EntrySelector uint16 `v8:"entrySelector"`
  RangeShift uint16 `v8:"rangeShift"`
  Table [1]uint16 `v8:"table"`
}

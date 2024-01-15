//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxOrderedListHeader
*/

type KerxOrderedListHeader struct {
  NPairs uint `v8:"nPairs"`
  SearchRange uint `v8:"searchRange"`
  EntrySelector uint `v8:"entrySelector"`
  RangeShift uint `v8:"rangeShift"`
  Table [1]uint `v8:"table"`
}

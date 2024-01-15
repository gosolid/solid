//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntDirectory
*/

type SfntDirectory struct {
  Format uint `v8:"format"`
  NumOffsets uint16 `v8:"numOffsets"`
  SearchRange uint16 `v8:"searchRange"`
  EntrySelector uint16 `v8:"entrySelector"`
  RangeShift uint16 `v8:"rangeShift"`
  Table [1]SfntDirectoryEntry `v8:"table"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntDirectoryEntry
*/

type SfntDirectoryEntry struct {
  TableTag uint `v8:"tableTag"`
  CheckSum uint `v8:"checkSum"`
  Offset uint `v8:"offset"`
  Length uint `v8:"length"`
}

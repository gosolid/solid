//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernVersion0SubtableHeader
*/

type KernVersion0SubtableHeader struct {
  Version uint16 `v8:"version"`
  Length uint16 `v8:"length"`
  StInfo uint16 `v8:"stInfo"`
  FsHeader void `v8:"fsHeader"`
}

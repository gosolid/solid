//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernSubtableHeader
*/

type KernSubtableHeader struct {
  Length int `v8:"length"`
  StInfo uint16 `v8:"stInfo"`
  TupleIndex int16 `v8:"tupleIndex"`
  FsHeader void `v8:"fsHeader"`
}

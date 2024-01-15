//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxSubtableHeader
*/

type KerxSubtableHeader struct {
  Length uint `v8:"length"`
  StInfo uint `v8:"stInfo"`
  TupleCount uint `v8:"tupleCount"`
  FsHeader void `v8:"fsHeader"`
}

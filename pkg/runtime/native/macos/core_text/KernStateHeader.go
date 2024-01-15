//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KernStateHeader
*/

type KernStateHeader struct {
  Header STHeader `v8:"header"`
  ValueTable uint16 `v8:"valueTable"`
  FirstTable [1]byte `v8:"firstTable"`
}

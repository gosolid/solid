//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxStateHeader
*/

type KerxStateHeader struct {
  Header STXHeader `v8:"header"`
  ValueTable uint `v8:"valueTable"`
  FirstTable [1]byte `v8:"firstTable"`
}

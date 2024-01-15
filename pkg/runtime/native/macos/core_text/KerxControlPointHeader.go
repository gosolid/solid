//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxControlPointHeader
*/

type KerxControlPointHeader struct {
  Header STXHeader `v8:"header"`
  Flags uint `v8:"flags"`
  FirstTable [1]byte `v8:"firstTable"`
}

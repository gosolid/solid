//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntCMapHeader
*/

type SfntCMapHeader struct {
  Version uint16 `v8:"version"`
  NumTables uint16 `v8:"numTables"`
  Encoding [1]SfntCMapEncoding `v8:"encoding"`
}

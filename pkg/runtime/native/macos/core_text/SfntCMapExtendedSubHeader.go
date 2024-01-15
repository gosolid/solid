//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntCMapExtendedSubHeader
*/

type SfntCMapExtendedSubHeader struct {
  Format uint16 `v8:"format"`
  Reserved uint16 `v8:"reserved"`
  Length uint `v8:"length"`
  Language uint `v8:"language"`
}

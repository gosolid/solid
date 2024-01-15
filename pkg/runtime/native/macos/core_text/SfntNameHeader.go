//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntNameHeader
*/

type SfntNameHeader struct {
  Format uint16 `v8:"format"`
  Count uint16 `v8:"count"`
  StringOffset uint16 `v8:"stringOffset"`
  Rec [1]SfntNameRecord `v8:"rec"`
}

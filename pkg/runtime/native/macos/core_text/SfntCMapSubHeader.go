//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntCMapSubHeader
*/

type SfntCMapSubHeader struct {
  Format uint16 `v8:"format"`
  Length uint16 `v8:"length"`
  LanguageID uint16 `v8:"languageID"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntFontDescriptor
*/

type SfntFontDescriptor struct {
  Name uint `v8:"name"`
  Value int `v8:"value"`
}

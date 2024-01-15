//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntFontFeatureSetting
*/

type SfntFontFeatureSetting struct {
  Setting uint16 `v8:"setting"`
  NameID int16 `v8:"nameID"`
}

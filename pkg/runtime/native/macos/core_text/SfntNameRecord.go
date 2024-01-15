//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntNameRecord
*/

type SfntNameRecord struct {
  PlatformID uint16 `v8:"platformID"`
  ScriptID uint16 `v8:"scriptID"`
  LanguageID uint16 `v8:"languageID"`
  NameID uint16 `v8:"nameID"`
  Length uint16 `v8:"length"`
  Offset uint16 `v8:"offset"`
}

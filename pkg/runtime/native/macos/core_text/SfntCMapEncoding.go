//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntCMapEncoding
*/

type SfntCMapEncoding struct {
  PlatformID uint16 `v8:"platformID"`
  ScriptID uint16 `v8:"scriptID"`
  Offset uint `v8:"offset"`
}

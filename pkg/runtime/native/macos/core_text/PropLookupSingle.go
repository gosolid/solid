//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.PropLookupSingle
*/

type PropLookupSingle struct {
  Glyph uint16 `v8:"glyph"`
  Props uint16 `v8:"props"`
}

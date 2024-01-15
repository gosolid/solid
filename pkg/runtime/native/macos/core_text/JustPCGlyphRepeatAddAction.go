//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustPCGlyphRepeatAddAction
*/

type JustPCGlyphRepeatAddAction struct {
  Flags uint16 `v8:"flags"`
  Glyph uint16 `v8:"glyph"`
}

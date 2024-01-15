//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.ROTAGlyphEntry
*/

type ROTAGlyphEntry struct {
  GlyphIndexOffset int16 `v8:"glyphIndexOffset"`
  HBaselineOffset int16 `v8:"hBaselineOffset"`
  VBaselineOffset int16 `v8:"vBaselineOffset"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.ALMXGlyphEntry
*/

type ALMXGlyphEntry struct {
  GlyphIndexOffset int16 `v8:"glyphIndexOffset"`
  HorizontalAdvance int16 `v8:"horizontalAdvance"`
  XOffsetToHOrigin int16 `v8:"xOffsetToHOrigin"`
  VerticalAdvance int16 `v8:"verticalAdvance"`
  YOffsetToVOrigin int16 `v8:"yOffsetToVOrigin"`
}

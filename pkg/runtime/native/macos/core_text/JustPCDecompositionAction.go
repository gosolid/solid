//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustPCDecompositionAction
*/

type JustPCDecompositionAction struct {
  LowerLimit int `v8:"lowerLimit"`
  UpperLimit int `v8:"upperLimit"`
  Order uint16 `v8:"order"`
  Count uint16 `v8:"count"`
  Glyphs [1]uint16 `v8:"glyphs"`
}

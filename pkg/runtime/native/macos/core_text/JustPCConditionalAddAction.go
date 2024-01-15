//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustPCConditionalAddAction
*/

type JustPCConditionalAddAction struct {
  SubstThreshold int `v8:"substThreshold"`
  AddGlyph uint16 `v8:"addGlyph"`
  SubstGlyph uint16 `v8:"substGlyph"`
}

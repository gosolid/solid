//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxAnchorPointAction
*/

type KerxAnchorPointAction struct {
  MarkAnchorPoint uint16 `v8:"markAnchorPoint"`
  CurrAnchorPoint uint16 `v8:"currAnchorPoint"`
}

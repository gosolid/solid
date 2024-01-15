//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustPCDuctilityAction
*/

type JustPCDuctilityAction struct {
  DuctilityAxis uint `v8:"ductilityAxis"`
  MinimumLimit int `v8:"minimumLimit"`
  NoStretchValue int `v8:"noStretchValue"`
  MaximumLimit int `v8:"maximumLimit"`
}

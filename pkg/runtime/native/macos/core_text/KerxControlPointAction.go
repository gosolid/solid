//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxControlPointAction
*/

type KerxControlPointAction struct {
  MarkControlPoint uint16 `v8:"markControlPoint"`
  CurrControlPoint uint16 `v8:"currControlPoint"`
}

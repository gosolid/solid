//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.KerxCoordinateAction
*/

type KerxCoordinateAction struct {
  MarkX uint16 `v8:"markX"`
  MarkY uint16 `v8:"markY"`
  CurrX uint16 `v8:"currX"`
  CurrY uint16 `v8:"currY"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.AnchorPoint
*/

type AnchorPoint struct {
  X int16 `v8:"x"`
  Y int16 `v8:"y"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.LtagStringRange
*/

type LtagStringRange struct {
  Offset uint16 `v8:"offset"`
  Length uint16 `v8:"length"`
}

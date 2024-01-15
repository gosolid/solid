//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustPCActionSubrecord
*/

type JustPCActionSubrecord struct {
  TheClass uint16 `v8:"theClass"`
  TheType uint16 `v8:"theType"`
  Length uint `v8:"length"`
  Data uint `v8:"data"`
}

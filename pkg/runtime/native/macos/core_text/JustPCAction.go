//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustPCAction
*/

type JustPCAction struct {
  ActionCount uint `v8:"actionCount"`
  Actions [1]JustPCActionSubrecord `v8:"actions"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustWidthDeltaGroup
*/

type JustWidthDeltaGroup struct {
  Count uint `v8:"count"`
  Entries [1]JustWidthDeltaEntry `v8:"entries"`
}

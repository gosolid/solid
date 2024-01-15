//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.LtagTable
*/

type LtagTable struct {
  Version uint `v8:"version"`
  Flags uint `v8:"flags"`
  NumTags uint `v8:"numTags"`
  TagRange [1]LtagStringRange `v8:"tagRange"`
}

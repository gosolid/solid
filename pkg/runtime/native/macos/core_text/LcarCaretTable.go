//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.LcarCaretTable
*/

type LcarCaretTable struct {
  Version int `v8:"version"`
  Format uint16 `v8:"format"`
  Lookup SFNTLookupTable `v8:"lookup"`
}

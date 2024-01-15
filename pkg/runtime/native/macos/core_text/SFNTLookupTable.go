//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.SFNTLookupTable
*/

type SFNTLookupTable struct {
  Format uint16 `v8:"format"`
  FsHeader void `v8:"fsHeader"`
}

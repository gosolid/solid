//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.BslnTable
*/

type BslnTable struct {
  Version int `v8:"version"`
  Format uint16 `v8:"format"`
  DefaultBaseline uint16 `v8:"defaultBaseline"`
  Parts void `v8:"parts"`
}

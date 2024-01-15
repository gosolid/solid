//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.TrakTable
*/

type TrakTable struct {
  Version int `v8:"version"`
  Format uint16 `v8:"format"`
  HorizOffset uint16 `v8:"horizOffset"`
  VertOffset uint16 `v8:"vertOffset"`
}

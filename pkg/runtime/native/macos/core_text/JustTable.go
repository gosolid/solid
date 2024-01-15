//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustTable
*/

type JustTable struct {
  Version int `v8:"version"`
  Format uint16 `v8:"format"`
  HorizHeaderOffset uint16 `v8:"horizHeaderOffset"`
  VertHeaderOffset uint16 `v8:"vertHeaderOffset"`
}

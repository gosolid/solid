//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.CTParagraphStyleSetting
*/

type CTParagraphStyleSetting struct {
  Spec int `v8:"spec"`
  ValueSize uint64 `v8:"valueSize"`
  Value void `v8:"value"`
}

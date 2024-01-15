//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntVariationAxis
*/

type SfntVariationAxis struct {
  AxisTag uint `v8:"axisTag"`
  MinValue int `v8:"minValue"`
  DefaultValue int `v8:"defaultValue"`
  MaxValue int `v8:"maxValue"`
  Flags int16 `v8:"flags"`
  NameID int16 `v8:"nameID"`
}

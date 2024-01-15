//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.JustWidthDeltaEntry
*/

type JustWidthDeltaEntry struct {
  JustClass uint `v8:"justClass"`
  BeforeGrowLimit int `v8:"beforeGrowLimit"`
  BeforeShrinkLimit int `v8:"beforeShrinkLimit"`
  AfterGrowLimit int `v8:"afterGrowLimit"`
  AfterShrinkLimit int `v8:"afterShrinkLimit"`
  GrowFlags uint16 `v8:"growFlags"`
  ShrinkFlags uint16 `v8:"shrinkFlags"`
}

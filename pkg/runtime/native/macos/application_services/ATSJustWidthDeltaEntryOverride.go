//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSJustWidthDeltaEntryOverride
*/

type ATSJustWidthDeltaEntryOverride struct {
  BeforeGrowLimit int `v8:"beforeGrowLimit"`
  BeforeShrinkLimit int `v8:"beforeShrinkLimit"`
  AfterGrowLimit int `v8:"afterGrowLimit"`
  AfterShrinkLimit int `v8:"afterShrinkLimit"`
  GrowFlags uint16 `v8:"growFlags"`
  ShrinkFlags uint16 `v8:"shrinkFlags"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntFeatureName
*/

type SfntFeatureName struct {
  FeatureType uint16 `v8:"featureType"`
  SettingCount uint16 `v8:"settingCount"`
  OffsetToSettings int `v8:"offsetToSettings"`
  FeatureFlags uint16 `v8:"featureFlags"`
  NameID int16 `v8:"nameID"`
}

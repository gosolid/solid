//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntFeatureHeader
*/

type SfntFeatureHeader struct {
  Version int `v8:"version"`
  FeatureNameCount uint16 `v8:"featureNameCount"`
  FeatureSetCount uint16 `v8:"featureSetCount"`
  Reserved int `v8:"reserved"`
  Names [1]SfntFeatureName `v8:"names"`
  Settings [1]SfntFontFeatureSetting `v8:"settings"`
  Runs [1]SfntFontRunFeature `v8:"runs"`
}

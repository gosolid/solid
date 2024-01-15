//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.sfntFontRunFeature
*/

type SfntFontRunFeature struct {
  FeatureType uint16 `v8:"featureType"`
  Setting uint16 `v8:"setting"`
}

//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MortFeatureEntry
*/

type MortFeatureEntry struct {
  FeatureType uint16 `v8:"featureType"`
  FeatureSelector uint16 `v8:"featureSelector"`
  EnableFlags uint `v8:"enableFlags"`
  DisableFlags uint `v8:"disableFlags"`
}

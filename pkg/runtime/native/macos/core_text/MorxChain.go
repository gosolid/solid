//js:package native/macos/core-text
package core_text

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreText.MorxChain
*/

type MorxChain struct {
  DefaultFlags uint `v8:"defaultFlags"`
  Length uint `v8:"length"`
  NFeatures uint `v8:"nFeatures"`
  NSubtables uint `v8:"nSubtables"`
  FeatureEntries [1]MortFeatureEntry `v8:"featureEntries"`
}

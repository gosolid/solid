//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECInfo
*/

type TECInfo struct {
  Format uint16 `v8:"format"`
  TecVersion uint16 `v8:"tecVersion"`
  TecTextConverterFeatures uint `v8:"tecTextConverterFeatures"`
  TecUnicodeConverterFeatures uint `v8:"tecUnicodeConverterFeatures"`
  TecTextCommonFeatures uint `v8:"tecTextCommonFeatures"`
  TecTextEncodingsFolderName [32]byte `v8:"tecTextEncodingsFolderName"`
  TecExtensionFileName [32]byte `v8:"tecExtensionFileName"`
  TecLowestTEFileVersion uint16 `v8:"tecLowestTEFileVersion"`
  TecHighestTEFileVersion uint16 `v8:"tecHighestTEFileVersion"`
}

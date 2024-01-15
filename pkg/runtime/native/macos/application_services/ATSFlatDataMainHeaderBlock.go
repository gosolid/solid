//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSFlatDataMainHeaderBlock
*/

type ATSFlatDataMainHeaderBlock struct {
  Version uint `v8:"version"`
  SizeOfDataBlock uint `v8:"sizeOfDataBlock"`
  OffsetToTextLayouts uint `v8:"offsetToTextLayouts"`
  OffsetToStyleRuns uint `v8:"offsetToStyleRuns"`
  OffsetToStyleList uint `v8:"offsetToStyleList"`
}

//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.MMCDeviceInterface
*/

type MMCDeviceInterface struct {
  Reserved void `v8:"reserved"`
  QueryInterface func(...any) (any, error) `v8:"queryInterface"`
  AddRef func(...any) (any, error) `v8:"addRef"`
  Release func(...any) (any, error) `v8:"release"`
  Version uint16 `v8:"version"`
  Revision uint16 `v8:"revision"`
  Inquiry func(...any) (any, error) `v8:"inquiry"`
  TestUnitReady func(...any) (any, error) `v8:"testUnitReady"`
  GetPerformance func(...any) (any, error) `v8:"getPerformance"`
  GetConfiguration func(...any) (any, error) `v8:"getConfiguration"`
  ModeSense10 func(...any) (any, error) `v8:"modeSense10"`
  SetWriteParametersModePage func(...any) (any, error) `v8:"setWriteParametersModePage"`
  GetTrayState func(...any) (any, error) `v8:"getTrayState"`
  SetTrayState func(...any) (any, error) `v8:"setTrayState"`
  ReadTableOfContents func(...any) (any, error) `v8:"readTableOfContents"`
  ReadDiscInformation func(...any) (any, error) `v8:"readDiscInformation"`
  ReadTrackInformation func(...any) (any, error) `v8:"readTrackInformation"`
  ReadDVDStructure func(...any) (any, error) `v8:"readDVDStructure"`
  GetSCSITaskDeviceInterface func(...any) (any, error) `v8:"getSCSITaskDeviceInterface"`
  GetPerformanceV2 func(...any) (any, error) `v8:"getPerformanceV2"`
  SetCDSpeed func(...any) (any, error) `v8:"setCDSpeed"`
  ReadFormatCapacities func(...any) (any, error) `v8:"readFormatCapacities"`
  ReadDiscStructure func(...any) (any, error) `v8:"readDiscStructure"`
  ReadDiscInformationV2 func(...any) (any, error) `v8:"readDiscInformationV2"`
  ReadTrackInformationV2 func(...any) (any, error) `v8:"readTrackInformationV2"`
  SetStreaming func(...any) (any, error) `v8:"setStreaming"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultiFunctLutB2AType
*/

type CMMultiFunctLutB2AType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  InputChannels byte `v8:"inputChannels"`
  OutputChannels byte `v8:"outputChannels"`
  Reserved2 uint16 `v8:"reserved2"`
  OffsetBcurves uint `v8:"offsetBcurves"`
  OffsetMatrix uint `v8:"offsetMatrix"`
  OffsetMcurves uint `v8:"offsetMcurves"`
  OffsetCLUT uint `v8:"offsetCLUT"`
  OffsetAcurves uint `v8:"offsetAcurves"`
  Data [1]byte `v8:"data"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMLut16Type
*/

type CMLut16Type struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  InputChannels byte `v8:"inputChannels"`
  OutputChannels byte `v8:"outputChannels"`
  GridPoints byte `v8:"gridPoints"`
  Reserved2 byte `v8:"reserved2"`
  Matrix [3][3]int `v8:"matrix"`
  InputTableEntries uint16 `v8:"inputTableEntries"`
  OutputTableEntries uint16 `v8:"outputTableEntries"`
  InputTable [1]uint16 `v8:"inputTable"`
}

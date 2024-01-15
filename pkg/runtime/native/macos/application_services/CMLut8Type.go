//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMLut8Type
*/

type CMLut8Type struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  InputChannels byte `v8:"inputChannels"`
  OutputChannels byte `v8:"outputChannels"`
  GridPoints byte `v8:"gridPoints"`
  Reserved2 byte `v8:"reserved2"`
  Matrix [3][3]int `v8:"matrix"`
  InputTable [1]byte `v8:"inputTable"`
}

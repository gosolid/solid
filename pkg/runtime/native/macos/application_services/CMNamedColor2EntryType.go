//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMNamedColor2EntryType
*/

type CMNamedColor2EntryType struct {
  RootName [32]byte `v8:"rootName"`
  PCSColorCoords [3]uint16 `v8:"pCSColorCoords"`
  DeviceColorCoords [1]uint16 `v8:"deviceColorCoords"`
}

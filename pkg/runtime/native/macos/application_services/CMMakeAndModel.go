//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMakeAndModel
*/

type CMMakeAndModel struct {
  Manufacturer uint `v8:"manufacturer"`
  Model uint `v8:"model"`
  SerialNumber uint `v8:"serialNumber"`
  ManufactureDate uint `v8:"manufactureDate"`
  Reserved1 uint `v8:"reserved1"`
  Reserved2 uint `v8:"reserved2"`
  Reserved3 uint `v8:"reserved3"`
  Reserved4 uint `v8:"reserved4"`
}

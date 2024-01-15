//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMCurveType
*/

type CMCurveType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  CountValue uint `v8:"countValue"`
  Data [1]uint16 `v8:"data"`
}

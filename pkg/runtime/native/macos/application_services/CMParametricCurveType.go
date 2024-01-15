//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMParametricCurveType
*/

type CMParametricCurveType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  FunctionType uint16 `v8:"functionType"`
  Reserved2 uint16 `v8:"reserved2"`
  Value [1]int `v8:"value"`
}

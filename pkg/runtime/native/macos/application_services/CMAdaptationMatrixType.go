//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMAdaptationMatrixType
*/

type CMAdaptationMatrixType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  AdaptationMatrix [9]int `v8:"adaptationMatrix"`
}

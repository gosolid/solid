//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMVideoCardGammaType
*/

type CMVideoCardGammaType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Gamma CMVideoCardGamma `v8:"gamma"`
}

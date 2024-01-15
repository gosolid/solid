//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMViewingConditionsType
*/

type CMViewingConditionsType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Illuminant CMFixedXYZColor `v8:"illuminant"`
  Surround CMFixedXYZColor `v8:"surround"`
  StdIlluminant uint `v8:"stdIlluminant"`
}

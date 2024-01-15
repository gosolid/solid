//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMeasurementType
*/

type CMMeasurementType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  StandardObserver uint `v8:"standardObserver"`
  BackingXYZ CMFixedXYZColor `v8:"backingXYZ"`
  Geometry uint `v8:"geometry"`
  Flare uint `v8:"flare"`
  Illuminant uint `v8:"illuminant"`
}

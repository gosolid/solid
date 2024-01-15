//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMPS2CRDVMSizeType
*/

type CMPS2CRDVMSizeType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  Count uint `v8:"count"`
  IntentCRD [1]CMIntentCRDVMSize `v8:"intentCRD"`
}

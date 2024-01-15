//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMNativeDisplayInfoType
*/

type CMNativeDisplayInfoType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  NativeDisplayInfo CMNativeDisplayInfo `v8:"nativeDisplayInfo"`
}

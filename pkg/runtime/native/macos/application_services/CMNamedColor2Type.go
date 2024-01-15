//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMNamedColor2Type
*/

type CMNamedColor2Type struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  VendorFlag uint `v8:"vendorFlag"`
  Count uint `v8:"count"`
  DeviceChannelCount uint `v8:"deviceChannelCount"`
  PrefixName [32]byte `v8:"prefixName"`
  SuffixName [32]byte `v8:"suffixName"`
  Data [1]byte `v8:"data"`
}

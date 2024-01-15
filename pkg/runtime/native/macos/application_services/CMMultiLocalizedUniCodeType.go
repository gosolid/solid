//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMultiLocalizedUniCodeType
*/

type CMMultiLocalizedUniCodeType struct {
  TypeDescriptor uint `v8:"typeDescriptor"`
  Reserved uint `v8:"reserved"`
  EntryCount uint `v8:"entryCount"`
  EntrySize uint `v8:"entrySize"`
}

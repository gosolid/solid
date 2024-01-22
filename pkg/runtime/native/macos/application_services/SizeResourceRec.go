//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.SizeResourceRec
*/

type SizeResourceRec struct {
  Flags uint16 `v8:"flags"`
  PreferredHeapSize uint `v8:"preferredHeapSize"`
  MinimumHeapSize uint `v8:"minimumHeapSize"`
}
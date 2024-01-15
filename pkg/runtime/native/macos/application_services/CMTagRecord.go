//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMTagRecord
*/

type CMTagRecord struct {
  Tag uint `v8:"tag"`
  ElementOffset uint `v8:"elementOffset"`
  ElementSize uint `v8:"elementSize"`
}

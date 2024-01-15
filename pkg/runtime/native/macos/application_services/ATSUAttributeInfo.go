//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ATSUAttributeInfo
*/

type ATSUAttributeInfo struct {
  FTag uint `v8:"fTag"`
  FValueSize uint64 `v8:"fValueSize"`
}

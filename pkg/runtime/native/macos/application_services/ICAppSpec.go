//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICAppSpec
*/

type ICAppSpec struct {
  FCreator uint `v8:"fCreator"`
  Name [64]byte `v8:"name"`
}

//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.ICAppSpecList
*/

type ICAppSpecList struct {
  NumberOfItems int16 `v8:"numberOfItems"`
  AppSpecs [1]ICAppSpec `v8:"appSpecs"`
}

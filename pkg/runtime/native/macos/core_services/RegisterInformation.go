//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.RegisterInformation
*/

type RegisterInformation struct {
  UnusedRegisterInformationField void `v8:"unusedRegisterInformationField"`
}

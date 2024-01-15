//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FPUInformation
*/

type FPUInformation struct {
  UnusedFPUInformationField void `v8:"unusedFPUInformationField"`
}

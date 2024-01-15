//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MachineInformation
*/

type MachineInformation struct {
  UnusedMachineInformationField void `v8:"unusedMachineInformationField"`
}

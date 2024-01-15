//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentDescription
*/

type ComponentDescription struct {
  ComponentType uint `v8:"componentType"`
  ComponentSubType uint `v8:"componentSubType"`
  ComponentManufacturer uint `v8:"componentManufacturer"`
  ComponentFlags uint `v8:"componentFlags"`
  ComponentFlagsMask uint `v8:"componentFlagsMask"`
}

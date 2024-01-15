//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentResource
*/

type ComponentResource struct {
  Cd ComponentDescription `v8:"cd"`
  Component ResourceSpec `v8:"component"`
  ComponentName ResourceSpec `v8:"componentName"`
  ComponentInfo ResourceSpec `v8:"componentInfo"`
  ComponentIcon ResourceSpec `v8:"componentIcon"`
}

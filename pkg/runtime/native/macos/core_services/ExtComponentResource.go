//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ExtComponentResource
*/

type ExtComponentResource struct {
  Cd ComponentDescription `v8:"cd"`
  Component ResourceSpec `v8:"component"`
  ComponentName ResourceSpec `v8:"componentName"`
  ComponentInfo ResourceSpec `v8:"componentInfo"`
  ComponentIcon ResourceSpec `v8:"componentIcon"`
  ComponentVersion int `v8:"componentVersion"`
  ComponentRegisterFlags int `v8:"componentRegisterFlags"`
  ComponentIconFamily int16 `v8:"componentIconFamily"`
  Count int `v8:"count"`
  PlatformArray [1]ComponentPlatformInfo `v8:"platformArray"`
}

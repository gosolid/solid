//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentResourceExtension
*/

type ComponentResourceExtension struct {
  ComponentVersion int `v8:"componentVersion"`
  ComponentRegisterFlags int `v8:"componentRegisterFlags"`
  ComponentIconFamily int16 `v8:"componentIconFamily"`
}

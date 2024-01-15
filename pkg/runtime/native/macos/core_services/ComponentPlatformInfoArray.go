//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentPlatformInfoArray
*/

type ComponentPlatformInfoArray struct {
  Count int `v8:"count"`
  PlatformArray [1]ComponentPlatformInfo `v8:"platformArray"`
}

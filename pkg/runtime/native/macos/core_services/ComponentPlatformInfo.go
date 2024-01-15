//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ComponentPlatformInfo
*/

type ComponentPlatformInfo struct {
  ComponentFlags int `v8:"componentFlags"`
  Component ResourceSpec `v8:"component"`
  PlatformType int16 `v8:"platformType"`
}

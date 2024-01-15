//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.ResourceSpec
*/

type ResourceSpec struct {
  ResType uint `v8:"resType"`
  ResID int16 `v8:"resID"`
}

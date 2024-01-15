//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.FSSpec
*/

type FSSpec struct {
  Hidden [70]byte `v8:"hidden"`
}

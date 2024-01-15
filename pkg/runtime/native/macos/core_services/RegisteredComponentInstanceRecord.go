//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.RegisteredComponentInstanceRecord
*/

type RegisteredComponentInstanceRecord struct {
  Data [1]int64 `v8:"data"`
}

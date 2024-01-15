//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AliasRecord
*/

type AliasRecord struct {
  Hidden [6]byte `v8:"hidden"`
}

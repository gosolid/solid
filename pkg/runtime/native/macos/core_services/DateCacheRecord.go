//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.DateCacheRecord
*/

type DateCacheRecord struct {
  Hidden [256]int16 `v8:"hidden"`
}

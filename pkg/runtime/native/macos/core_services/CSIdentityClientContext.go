//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.CSIdentityClientContext
*/

type CSIdentityClientContext struct {
  Version int64 `v8:"version"`
  Info void `v8:"info"`
  Retain *func(...any) (any, error) `v8:"retain"`
  Release *func(...any) (any, error) `v8:"release"`
  CopyDescription *func(...any) (any, error) `v8:"copyDescription"`
  StatusUpdated *func(...any) (any, error) `v8:"statusUpdated"`
}

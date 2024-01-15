//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.CSIdentityQueryClientContext
*/

type CSIdentityQueryClientContext struct {
  Version int64 `v8:"version"`
  Info void `v8:"info"`
  RetainInfo *func(...any) (any, error) `v8:"retainInfo"`
  ReleaseInfo *func(...any) (any, error) `v8:"releaseInfo"`
  CopyInfoDescription *func(...any) (any, error) `v8:"copyInfoDescription"`
  ReceiveEvent *func(...any) (any, error) `v8:"receiveEvent"`
}

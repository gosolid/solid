//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineFetchChangesContext : objc.NSObject
*/

type CKSyncEngineFetchChangesContext struct {
  *objc.NSObject

}

func (r *CKSyncEngineFetchChangesContext) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesContext) Reason() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesContext) Options() (*CKSyncEngineFetchChangesOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesContext) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


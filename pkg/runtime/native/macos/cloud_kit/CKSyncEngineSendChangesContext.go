//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineSendChangesContext : objc.NSObject
*/

type CKSyncEngineSendChangesContext struct {
  *objc.NSObject

}

func (r *CKSyncEngineSendChangesContext) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesContext) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesContext) Reason() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesContext) Options() (*CKSyncEngineSendChangesOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}


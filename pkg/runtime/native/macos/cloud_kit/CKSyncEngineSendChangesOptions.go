//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineSendChangesOptions : objc.NSObject
*/

type CKSyncEngineSendChangesOptions struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *CKSyncEngineSendChangesOptions) Scope() (*CKSyncEngineSendChangesScope, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesOptions) SetScope() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesOptions) OperationGroup() (*CKOperationGroup, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesOptions) SetOperationGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineSendChangesOptions) InitWithScope() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


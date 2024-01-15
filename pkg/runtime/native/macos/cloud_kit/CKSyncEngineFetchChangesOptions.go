//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKSyncEngineFetchChangesOptions : objc.NSObject
*/

type CKSyncEngineFetchChangesOptions struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *CKSyncEngineFetchChangesOptions) SetScope() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesOptions) OperationGroup() (*CKOperationGroup, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesOptions) SetOperationGroup() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesOptions) PrioritizedZoneIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesOptions) SetPrioritizedZoneIDs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesOptions) InitWithScope() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchChangesOptions) Scope() (*CKSyncEngineFetchChangesScope, error) {
  return nil, fmt.Errorf("unimplemented")
}


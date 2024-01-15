//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineFetchedDatabaseChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineFetchedDatabaseChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineFetchedDatabaseChangesEvent) Modifications() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedDatabaseChangesEvent) Deletions() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


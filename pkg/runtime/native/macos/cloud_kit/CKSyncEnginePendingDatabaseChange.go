//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CloudKit.CKSyncEnginePendingDatabaseChange : objc.NSObject
*/

type CKSyncEnginePendingDatabaseChange struct {
  *objc.NSObject

}

func (r *CKSyncEnginePendingDatabaseChange) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingDatabaseChange) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingDatabaseChange) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEnginePendingDatabaseChange) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


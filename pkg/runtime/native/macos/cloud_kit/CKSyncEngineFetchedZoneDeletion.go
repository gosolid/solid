//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKSyncEngineFetchedZoneDeletion : objc.NSObject
*/

type CKSyncEngineFetchedZoneDeletion struct {
  *objc.NSObject

}

func (r *CKSyncEngineFetchedZoneDeletion) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedZoneDeletion) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedZoneDeletion) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedZoneDeletion) Reason() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


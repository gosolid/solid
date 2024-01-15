//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineFetchedRecordDeletion : objc.NSObject
*/

type CKSyncEngineFetchedRecordDeletion struct {
  *objc.NSObject

}

func (r *CKSyncEngineFetchedRecordDeletion) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedRecordDeletion) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedRecordDeletion) RecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFetchedRecordDeletion) RecordType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKSyncEngineFailedRecordSave : objc.NSObject
*/

type CKSyncEngineFailedRecordSave struct {
  *objc.NSObject

}

func (r *CKSyncEngineFailedRecordSave) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFailedRecordSave) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFailedRecordSave) Record() (*CKRecord, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFailedRecordSave) Error() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}


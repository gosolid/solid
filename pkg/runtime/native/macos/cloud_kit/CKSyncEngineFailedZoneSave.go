//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineFailedZoneSave : objc.NSObject
*/

type CKSyncEngineFailedZoneSave struct {
  *objc.NSObject

}

func (r *CKSyncEngineFailedZoneSave) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFailedZoneSave) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFailedZoneSave) RecordZone() (*CKRecordZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineFailedZoneSave) Error() (*foundation.NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}


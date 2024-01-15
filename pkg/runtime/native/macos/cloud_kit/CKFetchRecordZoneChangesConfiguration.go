//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKFetchRecordZoneChangesConfiguration : objc.NSObject
*/

type CKFetchRecordZoneChangesConfiguration struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKFetchRecordZoneChangesConfiguration) SetResultsLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesConfiguration) DesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesConfiguration) SetDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesConfiguration) PreviousServerChangeToken() (*CKServerChangeToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesConfiguration) SetPreviousServerChangeToken() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesConfiguration) ResultsLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}


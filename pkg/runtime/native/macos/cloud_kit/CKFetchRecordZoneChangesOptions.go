//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKFetchRecordZoneChangesOptions : objc.NSObject
*/

type CKFetchRecordZoneChangesOptions struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKFetchRecordZoneChangesOptions) PreviousServerChangeToken() (*CKServerChangeToken, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOptions) SetPreviousServerChangeToken() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOptions) ResultsLimit() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOptions) SetResultsLimit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOptions) DesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchRecordZoneChangesOptions) SetDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}


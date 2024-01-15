//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKRecordZoneNotification : CloudKit.CKNotification
*/

type CKRecordZoneNotification struct {
  *CKNotification

}

func (r *CKRecordZoneNotification) RecordZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecordZoneNotification) DatabaseScope() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


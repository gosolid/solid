//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKQueryNotification : CloudKit.CKNotification
*/

type CKQueryNotification struct {
  *CKNotification

}

func (r *CKQueryNotification) RecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQueryNotification) DatabaseScope() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKQueryNotification) QueryNotificationReason() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKQueryNotification) RecordFields() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}


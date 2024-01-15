//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKDatabaseNotification : CloudKit.CKNotification
*/

type CKDatabaseNotification struct {
  *CKNotification

}

func (r *CKDatabaseNotification) DatabaseScope() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


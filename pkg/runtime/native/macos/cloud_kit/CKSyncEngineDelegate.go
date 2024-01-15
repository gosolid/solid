//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CloudKit.CKSyncEngineDelegate
*/

type CKSyncEngineDelegate struct {

}

func (r *CKSyncEngineDelegate) SyncEngine() error {
  return fmt.Errorf("unimplemented")
}


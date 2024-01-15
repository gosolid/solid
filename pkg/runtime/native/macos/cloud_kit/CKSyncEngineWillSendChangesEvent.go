//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEngineWillSendChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineWillSendChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineWillSendChangesEvent) Context() (*CKSyncEngineSendChangesContext, error) {
  return nil, fmt.Errorf("unimplemented")
}


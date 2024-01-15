//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEngineDidSendChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineDidSendChangesEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineDidSendChangesEvent) Context() (*CKSyncEngineSendChangesContext, error) {
  return nil, fmt.Errorf("unimplemented")
}


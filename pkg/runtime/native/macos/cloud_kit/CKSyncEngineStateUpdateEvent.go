//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKSyncEngineStateUpdateEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineStateUpdateEvent struct {
  *CKSyncEngineEvent

}

func (r *CKSyncEngineStateUpdateEvent) StateSerialization() (*CKSyncEngineStateSerialization, error) {
  return nil, fmt.Errorf("unimplemented")
}


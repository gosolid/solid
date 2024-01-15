//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKSyncEngineEvent : objc.NSObject
*/

type CKSyncEngineEvent struct {
  *objc.NSObject

}

func (r *CKSyncEngineEvent) FetchedRecordZoneChangesEvent() (*CKSyncEngineFetchedRecordZoneChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) DidFetchChangesEvent() (*CKSyncEngineDidFetchChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) DidSendChangesEvent() (*CKSyncEngineDidSendChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) WillFetchRecordZoneChangesEvent() (*CKSyncEngineWillFetchRecordZoneChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) WillSendChangesEvent() (*CKSyncEngineWillSendChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) SentRecordZoneChangesEvent() (*CKSyncEngineSentRecordZoneChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) DidFetchRecordZoneChangesEvent() (*CKSyncEngineDidFetchRecordZoneChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) FetchedDatabaseChangesEvent() (*CKSyncEngineFetchedDatabaseChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) StateUpdateEvent() (*CKSyncEngineStateUpdateEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) AccountChangeEvent() (*CKSyncEngineAccountChangeEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) SentDatabaseChangesEvent() (*CKSyncEngineSentDatabaseChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) WillFetchChangesEvent() (*CKSyncEngineWillFetchChangesEvent, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineEvent) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


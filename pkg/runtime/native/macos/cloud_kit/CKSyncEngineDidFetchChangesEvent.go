//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
interface CloudKit.CKSyncEngineDidFetchChangesEvent : CloudKit.CKSyncEngineEvent
*/

type CKSyncEngineDidFetchChangesEvent struct {
  *CKSyncEngineEvent

}


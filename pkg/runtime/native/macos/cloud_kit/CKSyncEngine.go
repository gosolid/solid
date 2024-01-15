//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKSyncEngine : objc.NSObject
*/

type CKSyncEngine struct {
  *objc.NSObject

}

func (r *CKSyncEngine) FetchChangesWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) SendChangesWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) CancelOperationsWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) FetchChangesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) SendChangesWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) Database() (*CKDatabase, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) State() (*CKSyncEngineState, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) InitWithConfiguration() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngine) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


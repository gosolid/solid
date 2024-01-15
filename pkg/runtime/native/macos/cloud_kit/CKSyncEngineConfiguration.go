//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKSyncEngineConfiguration : objc.NSObject
*/

type CKSyncEngineConfiguration struct {
  *objc.NSObject

}

func (r *CKSyncEngineConfiguration) InitWithDatabase() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) SetDatabase() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) StateSerialization() (*CKSyncEngineStateSerialization, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) SubscriptionID() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) AutomaticallySync() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) SetAutomaticallySync() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) Database() (*CKDatabase, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) SetStateSerialization() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSyncEngineConfiguration) SetSubscriptionID() error {
  return fmt.Errorf("unimplemented")
}


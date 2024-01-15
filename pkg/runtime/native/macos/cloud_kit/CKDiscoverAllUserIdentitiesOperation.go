//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKDiscoverAllUserIdentitiesOperation : CloudKit.CKOperation
*/

type CKDiscoverAllUserIdentitiesOperation struct {
  *CKOperation

}

func (r *CKDiscoverAllUserIdentitiesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverAllUserIdentitiesOperation) UserIdentityDiscoveredBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverAllUserIdentitiesOperation) SetUserIdentityDiscoveredBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDiscoverAllUserIdentitiesOperation) DiscoverAllUserIdentitiesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverAllUserIdentitiesOperation) SetDiscoverAllUserIdentitiesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}


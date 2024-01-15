//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKDiscoverUserIdentitiesOperation : CloudKit.CKOperation
*/

type CKDiscoverUserIdentitiesOperation struct {
  *CKOperation

}

func (r *CKDiscoverUserIdentitiesOperation) SetUserIdentityLookupInfos() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) UserIdentityDiscoveredBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) SetUserIdentityDiscoveredBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) DiscoverUserIdentitiesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) SetDiscoverUserIdentitiesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) InitWithUserIdentityLookupInfos() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDiscoverUserIdentitiesOperation) UserIdentityLookupInfos() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


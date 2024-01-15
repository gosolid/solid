//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKModifySubscriptionsOperation : CloudKit.CKDatabaseOperation
*/

type CKModifySubscriptionsOperation struct {
  *CKDatabaseOperation

}

func (r *CKModifySubscriptionsOperation) PerSubscriptionSaveBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SetPerSubscriptionSaveBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SetPerSubscriptionDeleteBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SubscriptionsToSave() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SetSubscriptionsToSave() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SubscriptionIDsToDelete() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SetSubscriptionIDsToDelete() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) SetModifySubscriptionsCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) InitWithSubscriptionsToSave() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) PerSubscriptionDeleteBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifySubscriptionsOperation) ModifySubscriptionsCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


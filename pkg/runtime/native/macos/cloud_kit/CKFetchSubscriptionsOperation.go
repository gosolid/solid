//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchSubscriptionsOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchSubscriptionsOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchSubscriptionsOperation) InitWithSubscriptionIDs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) SubscriptionIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) SetSubscriptionIDs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) PerSubscriptionCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) SetPerSubscriptionCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) FetchSubscriptionCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) SetFetchSubscriptionCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchSubscriptionsOperation) FetchAllSubscriptionsOperation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


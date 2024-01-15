//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchWebAuthTokenOperation : CloudKit.CKDatabaseOperation
*/

type CKFetchWebAuthTokenOperation struct {
  *CKDatabaseOperation

}

func (r *CKFetchWebAuthTokenOperation) APIToken() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchWebAuthTokenOperation) SetAPIToken() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchWebAuthTokenOperation) FetchWebAuthTokenCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchWebAuthTokenOperation) SetFetchWebAuthTokenCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchWebAuthTokenOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchWebAuthTokenOperation) InitWithAPIToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


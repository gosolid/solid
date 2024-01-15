//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKAcceptSharesOperation : CloudKit.CKOperation
*/

type CKAcceptSharesOperation struct {
  *CKOperation

}

func (r *CKAcceptSharesOperation) ShareMetadatas() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) SetShareMetadatas() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) PerShareCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) SetPerShareCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) AcceptSharesCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) SetAcceptSharesCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAcceptSharesOperation) InitWithShareMetadatas() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


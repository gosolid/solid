//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchShareMetadataOperation : CloudKit.CKOperation
*/

type CKFetchShareMetadataOperation struct {
  *CKOperation

}

func (r *CKFetchShareMetadataOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) ShareURLs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) SetShouldFetchRootRecord() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) PerShareMetadataBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) SetFetchShareMetadataCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) InitWithShareURLs() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) SetShareURLs() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) ShouldFetchRootRecord() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) RootRecordDesiredKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) SetRootRecordDesiredKeys() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) SetPerShareMetadataBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareMetadataOperation) FetchShareMetadataCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


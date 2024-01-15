//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CloudKit.CKModifyBadgeOperation : CloudKit.CKOperation
*/

type CKModifyBadgeOperation struct {
  *CKOperation

}

func (r *CKModifyBadgeOperation) ModifyBadgeCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyBadgeOperation) SetModifyBadgeCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKModifyBadgeOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyBadgeOperation) InitWithBadgeValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKModifyBadgeOperation) BadgeValue() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKModifyBadgeOperation) SetBadgeValue() error {
  return fmt.Errorf("unimplemented")
}


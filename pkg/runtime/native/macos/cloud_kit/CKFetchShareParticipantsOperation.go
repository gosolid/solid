//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKFetchShareParticipantsOperation : CloudKit.CKOperation
*/

type CKFetchShareParticipantsOperation struct {
  *CKOperation

}

func (r *CKFetchShareParticipantsOperation) SetFetchShareParticipantsCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) UserIdentityLookupInfos() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) SetPerShareParticipantCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) SetShareParticipantFetchedBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) PerShareParticipantCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) FetchShareParticipantsCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) InitWithUserIdentityLookupInfos() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) SetUserIdentityLookupInfos() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKFetchShareParticipantsOperation) ShareParticipantFetchedBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}


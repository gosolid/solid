//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKShare : CloudKit.CKRecord
*/

type CKShare struct {
  *CKRecord
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKShare) RemoveParticipant() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKShare) PublicPermission() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShare) Owner() (*CKShareParticipant, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) InitWithRootRecord() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) InitWithRecordZoneID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) AddParticipant() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKShare) Participants() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) CurrentUserParticipant() (*CKShareParticipant, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) InitWithRecordType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShare) SetPublicPermission() error {
  return fmt.Errorf("unimplemented")
}


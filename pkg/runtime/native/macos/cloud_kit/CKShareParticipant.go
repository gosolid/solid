//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKShareParticipant : objc.NSObject
*/

type CKShareParticipant struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKShareParticipant) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) Permission() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) Role() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) SetRole() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) AcceptanceStatus() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) SetPermission() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareParticipant) UserIdentity() (*CKUserIdentity, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKAllowedSharingOptions : objc.NSObject
*/

type CKAllowedSharingOptions struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKAllowedSharingOptions) SetAllowedParticipantPermissionOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKAllowedSharingOptions) AllowedParticipantAccessOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKAllowedSharingOptions) SetAllowedParticipantAccessOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKAllowedSharingOptions) StandardOptions() (*CKAllowedSharingOptions, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAllowedSharingOptions) InitWithAllowedParticipantPermissionOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKAllowedSharingOptions) AllowedParticipantPermissionOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


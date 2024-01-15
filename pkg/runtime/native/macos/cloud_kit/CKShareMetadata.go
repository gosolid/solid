//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKShareMetadata : objc.NSObject
*/

type CKShareMetadata struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CKShareMetadata) ParticipantRole() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) ParticipantStatus() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) OwnerIdentity() (*CKUserIdentity, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) RootRecord() (*CKRecord, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) ContainerIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) Share() (*CKShare, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) HierarchicalRootRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) RootRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) ParticipantPermission() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKShareMetadata) ParticipantType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKUserIdentity : objc.NSObject
*/

type CKUserIdentity struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKUserIdentity) LookupInfo() (*CKUserIdentityLookupInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentity) NameComponents() (*foundation.NSPersonNameComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentity) UserRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentity) ContactIdentifiers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentity) HasiCloudAccount() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentity) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentity) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


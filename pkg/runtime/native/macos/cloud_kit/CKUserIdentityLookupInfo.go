//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKUserIdentityLookupInfo : objc.NSObject
*/

type CKUserIdentityLookupInfo struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKUserIdentityLookupInfo) PhoneNumber() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) InitWithEmailAddress() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) InitWithUserRecordID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) LookupInfosWithPhoneNumbers() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) LookupInfosWithRecordIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) EmailAddress() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) UserRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) InitWithPhoneNumber() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKUserIdentityLookupInfo) LookupInfosWithEmails() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


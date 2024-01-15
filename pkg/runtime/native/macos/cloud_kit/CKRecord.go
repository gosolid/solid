//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKRecord : objc.NSObject
*/

type CKRecord struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKRecord) InitWithRecordType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) ObjectForKeyedSubscript() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) SetParentReferenceFromRecord() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKRecord) RecordType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) CreationDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) Parent() (*CKReference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) RecordChangeTag() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) CreatorUserRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) LastModifiedUserRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) ModificationDate() (*foundation.NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) SetParent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKRecord) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) ChangedKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) EncodeSystemFieldsWithCoder() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKRecord) SetParentReferenceFromRecordID() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKRecord) Share() (*CKReference, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKRecord) AllKeys() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) AllTokens() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKRecord) RecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}


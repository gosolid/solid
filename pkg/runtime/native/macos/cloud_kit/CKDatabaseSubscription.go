//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKDatabaseSubscription : CloudKit.CKSubscription
*/

type CKDatabaseSubscription struct {
  *CKSubscription
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKDatabaseSubscription) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDatabaseSubscription) RecordType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDatabaseSubscription) SetRecordType() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKDatabaseSubscription) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDatabaseSubscription) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKDatabaseSubscription) InitWithSubscriptionID() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


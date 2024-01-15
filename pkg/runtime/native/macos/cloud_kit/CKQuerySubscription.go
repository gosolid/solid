//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKQuerySubscription : CloudKit.CKSubscription
*/

type CKQuerySubscription struct {
  *CKSubscription
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKQuerySubscription) InitWithRecordType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuerySubscription) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuerySubscription) RecordType() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuerySubscription) Predicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuerySubscription) ZoneID() (*CKRecordZoneID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKQuerySubscription) SetZoneID() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKQuerySubscription) QuerySubscriptionOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


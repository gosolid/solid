//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CloudKit.CKSubscription : objc.NSObject
*/

type CKSubscription struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CKSubscription) NotificationInfo() (*CKNotificationInfo, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSubscription) SetNotificationInfo() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKSubscription) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSubscription) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSubscription) SubscriptionID() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKSubscription) SubscriptionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


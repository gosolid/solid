//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CloudKit.CKNotification : objc.NSObject
*/

type CKNotification struct {
  *objc.NSObject

}

func (r *CKNotification) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotification) NotificationFromRemoteNotificationDictionary() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotification) NotificationType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CKNotification) NotificationID() (*CKNotificationID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotification) IsPruned() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CKNotification) SubscriptionID() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotification) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotification) SubscriptionOwnerUserRecordID() (*CKRecordID, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKNotification) ContainerIdentifier() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


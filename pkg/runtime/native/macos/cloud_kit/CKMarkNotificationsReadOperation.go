//js:package native/macos/cloud-kit
package cloud_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CloudKit.CKMarkNotificationsReadOperation : CloudKit.CKOperation
*/

type CKMarkNotificationsReadOperation struct {
  *CKOperation

}

func (r *CKMarkNotificationsReadOperation) MarkNotificationsReadCompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKMarkNotificationsReadOperation) SetMarkNotificationsReadCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CKMarkNotificationsReadOperation) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKMarkNotificationsReadOperation) InitWithNotificationIDsToMarkRead() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKMarkNotificationsReadOperation) NotificationIDs() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CKMarkNotificationsReadOperation) SetNotificationIDs() error {
  return fmt.Errorf("unimplemented")
}


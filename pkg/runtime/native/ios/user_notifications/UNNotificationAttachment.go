//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationAttachment : objc.NSObject
*/

type UNNotificationAttachment struct {
  *objc.NSObject

}

func (r *UNNotificationAttachment) AttachmentWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAttachment) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAttachment) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAttachment) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAttachment) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


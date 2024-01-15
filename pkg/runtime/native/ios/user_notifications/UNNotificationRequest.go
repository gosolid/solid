//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationRequest : objc.NSObject
*/

type UNNotificationRequest struct {
  *objc.NSObject

}

func (r *UNNotificationRequest) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationRequest) Content() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationRequest) Trigger() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationRequest) RequestWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationRequest) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


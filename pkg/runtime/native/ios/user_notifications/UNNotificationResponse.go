//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationResponse : objc.NSObject
*/

type UNNotificationResponse struct {
  *objc.NSObject

}

func (r *UNNotificationResponse) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationResponse) Notification() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationResponse) ActionIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationTrigger : objc.NSObject
*/

type UNNotificationTrigger struct {
  *objc.NSObject

}

func (r *UNNotificationTrigger) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationTrigger) Repeats() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


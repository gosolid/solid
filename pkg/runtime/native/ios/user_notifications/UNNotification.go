//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotification : objc.NSObject
*/

type UNNotification struct {
  *objc.NSObject

}

func (r *UNNotification) Date() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotification) Request() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotification) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


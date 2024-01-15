//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationAction : objc.NSObject
*/

type UNNotificationAction struct {
  *objc.NSObject

}

func (r *UNNotificationAction) ActionWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAction) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAction) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationAction) Icon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


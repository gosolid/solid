//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationActionIcon : objc.NSObject
*/

type UNNotificationActionIcon struct {
  *objc.NSObject

}

func (r *UNNotificationActionIcon) IconWithTemplateImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationActionIcon) IconWithSystemImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationActionIcon) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

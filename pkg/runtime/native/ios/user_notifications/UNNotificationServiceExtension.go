//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationServiceExtension : objc.NSObject
*/

type UNNotificationServiceExtension struct {
  *objc.NSObject

}

func (r *UNNotificationServiceExtension) DidReceiveNotificationRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationServiceExtension) ServiceExtensionTimeWillExpire() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


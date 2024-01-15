//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationCategory : objc.NSObject
*/

type UNNotificationCategory struct {
  *objc.NSObject

}

func (r *UNNotificationCategory) HiddenPreviewsBodyPlaceholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) CategorySummaryFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) CategoryWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) Actions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) IntentIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationCategory) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


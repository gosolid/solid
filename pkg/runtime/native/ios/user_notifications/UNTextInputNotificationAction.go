//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UserNotifications.UNTextInputNotificationAction : UserNotifications.UNNotificationAction
*/

type UNTextInputNotificationAction struct {
  *UNNotificationAction

}

func (r *UNTextInputNotificationAction) TextInputButtonTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNTextInputNotificationAction) TextInputPlaceholder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNTextInputNotificationAction) ActionWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


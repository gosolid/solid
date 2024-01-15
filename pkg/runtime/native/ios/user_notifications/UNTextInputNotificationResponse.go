//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UserNotifications.UNTextInputNotificationResponse : UserNotifications.UNNotificationResponse
*/

type UNTextInputNotificationResponse struct {
  *UNNotificationResponse

}

func (r *UNTextInputNotificationResponse) UserText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


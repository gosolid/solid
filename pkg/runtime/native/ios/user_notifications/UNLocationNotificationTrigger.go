//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UserNotifications.UNLocationNotificationTrigger : UserNotifications.UNNotificationTrigger
*/

type UNLocationNotificationTrigger struct {
  *UNNotificationTrigger

}

func (r *UNLocationNotificationTrigger) TriggerWithRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNLocationNotificationTrigger) Region() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


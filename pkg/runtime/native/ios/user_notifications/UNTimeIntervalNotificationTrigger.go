//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UserNotifications.UNTimeIntervalNotificationTrigger : UserNotifications.UNNotificationTrigger
*/

type UNTimeIntervalNotificationTrigger struct {
  *UNNotificationTrigger

}

func (r *UNTimeIntervalNotificationTrigger) TriggerWithTimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNTimeIntervalNotificationTrigger) NextTriggerDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNTimeIntervalNotificationTrigger) TimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


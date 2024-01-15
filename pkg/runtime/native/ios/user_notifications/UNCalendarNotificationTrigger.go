//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UserNotifications.UNCalendarNotificationTrigger : UserNotifications.UNNotificationTrigger
*/

type UNCalendarNotificationTrigger struct {
  *UNNotificationTrigger

}

func (r *UNCalendarNotificationTrigger) TriggerWithDateMatchingComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNCalendarNotificationTrigger) NextTriggerDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNCalendarNotificationTrigger) DateComponents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


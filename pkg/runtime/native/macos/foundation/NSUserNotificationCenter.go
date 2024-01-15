//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUserNotificationCenter : objc.NSObject
*/

type NSUserNotificationCenter struct {
  *objc.NSObject

}

func (r *NSUserNotificationCenter) DeliverNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) DefaultUserNotificationCenter() (*NSUserNotificationCenter, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) ScheduledNotifications() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) ScheduleNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) RemoveScheduledNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) RemoveDeliveredNotification() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) RemoveAllDeliveredNotifications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) SetScheduledNotifications() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserNotificationCenter) DeliveredNotifications() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


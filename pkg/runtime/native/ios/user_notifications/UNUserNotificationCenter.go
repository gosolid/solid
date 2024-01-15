//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNUserNotificationCenter : objc.NSObject
*/

type UNUserNotificationCenter struct {
  *objc.NSObject

}

func (r *UNUserNotificationCenter) SetNotificationCategories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) GetNotificationCategoriesWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) GetPendingNotificationRequestsWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) RemovePendingNotificationRequestsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) GetDeliveredNotificationsWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) SetBadgeCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) CurrentNotificationCenter() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) AddNotificationRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) SupportsContentExtensions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) RequestAuthorizationWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) RemoveDeliveredNotificationsWithIdentifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) GetNotificationSettingsWithCompletionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) RemoveAllPendingNotificationRequests() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNUserNotificationCenter) RemoveAllDeliveredNotifications() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


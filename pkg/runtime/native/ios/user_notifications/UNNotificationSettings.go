//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UserNotifications.UNNotificationSettings : objc.NSObject
*/

type UNNotificationSettings struct {
  *objc.NSObject

}

func (r *UNNotificationSettings) TimeSensitiveSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) ScheduledDeliverySetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) DirectMessagesSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) AlertSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) NotificationCenterSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) AlertStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) CriticalAlertSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) AuthorizationStatus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) SoundSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) BadgeSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) CarPlaySetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) ShowPreviewsSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) LockScreenSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) ProvidesAppNotificationSettings() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationSettings) AnnouncementSetting() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


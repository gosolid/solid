//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UserNotifications.UNNotificationContent : objc.NSObject
*/

type UNNotificationContent struct {
  *objc.NSObject

}

func (r *UNNotificationContent) Body() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) ContentByUpdatingWithProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) ThreadIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) SummaryArgument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) SummaryArgumentCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) TargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) FilterCriteria() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) Badge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) CategoryIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) LaunchImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) Sound() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) RelevanceScore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) Attachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) InterruptionLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNNotificationContent) Subtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


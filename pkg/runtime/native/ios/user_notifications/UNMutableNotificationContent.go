//js:package native/ios/user-notifications
package user_notifications

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UserNotifications.UNMutableNotificationContent : UserNotifications.UNNotificationContent
*/

type UNMutableNotificationContent struct {
  *UNNotificationContent

}

func (r *UNMutableNotificationContent) SetRelevanceScore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) Body() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetBody() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetUserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) TargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) InterruptionLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetSubtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) ThreadIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) RelevanceScore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) FilterCriteria() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SummaryArgument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) CategoryIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) LaunchImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetSummaryArgumentCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetAttachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetCategoryIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetSound() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) Subtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetInterruptionLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) Attachments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) Badge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetBadge() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) Sound() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetFilterCriteria() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetLaunchImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetThreadIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetSummaryArgument() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SummaryArgumentCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UNMutableNotificationContent) SetTargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


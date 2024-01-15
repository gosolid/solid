//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIMutableUserNotificationCategory : UIKit.UIUserNotificationCategory
*/

type UIMutableUserNotificationCategory struct {
  *UIUserNotificationCategory

}

func (r *UIMutableUserNotificationCategory) SetActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationCategory) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationCategory) SetIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


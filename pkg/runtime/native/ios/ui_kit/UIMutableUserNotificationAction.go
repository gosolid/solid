//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIMutableUserNotificationAction : UIKit.UIUserNotificationAction
*/

type UIMutableUserNotificationAction struct {
  *UIUserNotificationAction

}

func (r *UIMutableUserNotificationAction) SetIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) SetTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) SetBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) Parameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) SetActivationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) Behavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) SetParameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) SetAuthenticationRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) IsDestructive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) SetDestructive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) ActivationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIMutableUserNotificationAction) IsAuthenticationRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


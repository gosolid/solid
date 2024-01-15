//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIUserNotificationAction : objc.NSObject
*/

type UIUserNotificationAction struct {
  *objc.NSObject

}

func (r *UIUserNotificationAction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) Title() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) ActivationMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) Behavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) Parameters() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) IsAuthenticationRequired() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationAction) IsDestructive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


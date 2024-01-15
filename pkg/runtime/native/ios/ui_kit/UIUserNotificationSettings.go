//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIUserNotificationSettings : objc.NSObject
*/

type UIUserNotificationSettings struct {
  *objc.NSObject

}

func (r *UIUserNotificationSettings) SettingsForTypes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationSettings) Types() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationSettings) Categories() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIUserNotificationCategory : objc.NSObject
*/

type UIUserNotificationCategory struct {
  *objc.NSObject

}

func (r *UIUserNotificationCategory) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationCategory) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationCategory) ActionsForContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIUserNotificationCategory) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


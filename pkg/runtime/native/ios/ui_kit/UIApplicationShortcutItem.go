//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIApplicationShortcutItem : objc.NSObject
*/

type UIApplicationShortcutItem struct {
  *objc.NSObject

}

func (r *UIApplicationShortcutItem) InitWithType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) LocalizedTitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) LocalizedSubtitle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) Icon() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) TargetContentIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutItem) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


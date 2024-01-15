//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIApplicationShortcutIcon : objc.NSObject
*/

type UIApplicationShortcutIcon struct {
  *objc.NSObject

}

func (r *UIApplicationShortcutIcon) IconWithTemplateImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutIcon) IconWithSystemImageName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIApplicationShortcutIcon) IconWithType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


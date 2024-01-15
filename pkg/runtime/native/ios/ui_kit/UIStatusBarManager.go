//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIStatusBarManager : objc.NSObject
*/

type UIStatusBarManager struct {
  *objc.NSObject

}

func (r *UIStatusBarManager) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStatusBarManager) StatusBarStyle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStatusBarManager) IsStatusBarHidden() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStatusBarManager) StatusBarFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIStatusBarManager) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


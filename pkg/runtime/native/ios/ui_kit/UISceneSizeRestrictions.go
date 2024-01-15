//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneSizeRestrictions : objc.NSObject
*/

type UISceneSizeRestrictions struct {
  *objc.NSObject

}

func (r *UISceneSizeRestrictions) SetAllowsFullScreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) MinimumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) SetMinimumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) MaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) SetMaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneSizeRestrictions) AllowsFullScreen() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


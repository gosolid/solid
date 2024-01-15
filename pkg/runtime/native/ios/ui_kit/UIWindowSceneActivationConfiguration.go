//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIWindowSceneActivationConfiguration : objc.NSObject
*/

type UIWindowSceneActivationConfiguration struct {
  *objc.NSObject

}

func (r *UIWindowSceneActivationConfiguration) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) SetOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) Preview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) SetPreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) InitWithUserActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationConfiguration) UserActivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


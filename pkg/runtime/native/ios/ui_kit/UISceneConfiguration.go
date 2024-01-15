//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISceneConfiguration : objc.NSObject
*/

type UISceneConfiguration struct {
  *objc.NSObject

}

func (r *UISceneConfiguration) Storyboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) SetStoryboard() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) Role() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) SetSceneClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) DelegateClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) SceneClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) SetDelegateClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) ConfigurationWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneConfiguration) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UISceneWindowingBehaviors : objc.NSObject
*/

type UISceneWindowingBehaviors struct {
  *objc.NSObject

}

func (r *UISceneWindowingBehaviors) IsMiniaturizable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneWindowingBehaviors) SetMiniaturizable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneWindowingBehaviors) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneWindowingBehaviors) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneWindowingBehaviors) IsClosable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISceneWindowingBehaviors) SetClosable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


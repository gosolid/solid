//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIWindowSceneActivationInteraction : objc.NSObject
*/

type UIWindowSceneActivationInteraction struct {
  *objc.NSObject

}

func (r *UIWindowSceneActivationInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationInteraction) InitWithConfigurationProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIWindowSceneActivationInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


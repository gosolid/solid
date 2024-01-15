//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UISpringLoadedInteraction : objc.NSObject
*/

type UISpringLoadedInteraction struct {
  *objc.NSObject

}

func (r *UISpringLoadedInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringLoadedInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringLoadedInteraction) InitWithInteractionBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringLoadedInteraction) InitWithActivationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringLoadedInteraction) InteractionBehavior() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UISpringLoadedInteraction) InteractionEffect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFocusMovementHint : objc.NSObject
*/

type UIFocusMovementHint struct {
  *objc.NSObject

}

func (r *UIFocusMovementHint) Translation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusMovementHint) InteractionTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusMovementHint) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusMovementHint) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusMovementHint) MovementDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusMovementHint) PerspectiveTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFocusMovementHint) Rotation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


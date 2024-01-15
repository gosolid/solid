//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPencilInteraction : objc.NSObject
*/

type UIPencilInteraction struct {
  *objc.NSObject

}

func (r *UIPencilInteraction) PreferredTapAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPencilInteraction) PrefersPencilOnlyDrawing() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPencilInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPencilInteraction) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPencilInteraction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPencilInteraction) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


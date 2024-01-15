//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDragInteraction : objc.NSObject
*/

type UIDragInteraction struct {
  *objc.NSObject

}

func (r *UIDragInteraction) SetAllowsSimultaneousRecognitionDuringLift() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) IsEnabledByDefault() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) AllowsSimultaneousRecognitionDuringLift() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDragInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


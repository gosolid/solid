//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIResponder : objc.NSObject
*/

type UIResponder struct {
  *objc.NSObject

}

func (r *UIResponder) TouchesEstimatedPropertiesUpdated() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) PressesCancelled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) MotionEnded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) BuildMenuWithBuilder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) TargetForAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) ValidateCommand() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) TouchesEnded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) RemoteControlReceivedWithEvent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) BecomeFirstResponder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) TouchesBegan() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) TouchesCancelled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) CanPerformAction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) UndoManager() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) PressesBegan() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) PressesChanged() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) NextResponder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) EditingInteractionConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) ResignFirstResponder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) MotionBegan() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) CanResignFirstResponder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) TouchesMoved() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) PressesEnded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) IsFirstResponder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) MotionCancelled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIResponder) CanBecomeFirstResponder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


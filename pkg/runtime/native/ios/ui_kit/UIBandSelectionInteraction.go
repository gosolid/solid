//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIBandSelectionInteraction : objc.NSObject
*/

type UIBandSelectionInteraction struct {
  *objc.NSObject

}

func (r *UIBandSelectionInteraction) InitWithSelectionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) SetEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) State() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) SelectionRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) InitialModifierFlags() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) IsEnabled() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) ShouldBeginHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIBandSelectionInteraction) SetShouldBeginHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


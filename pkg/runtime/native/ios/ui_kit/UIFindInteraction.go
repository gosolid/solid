//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIFindInteraction : objc.NSObject
*/

type UIFindInteraction struct {
  *objc.NSObject

}

func (r *UIFindInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) OptionsMenuProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) DismissFindNavigator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) SetSearchText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) SetOptionsMenuProvider() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) FindPrevious() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) UpdateResultCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) SearchText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) ReplacementText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) SetReplacementText() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) InitWithSessionDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) FindNext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) IsFindNavigatorVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) ActiveFindSession() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIFindInteraction) PresentFindNavigatorShowingReplace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


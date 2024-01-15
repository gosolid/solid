//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextFormattingCoordinator : objc.NSObject
*/

type UITextFormattingCoordinator struct {
  *objc.NSObject

}

func (r *UITextFormattingCoordinator) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) IsFontPanelVisible() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) TextFormattingCoordinatorForWindowScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) InitWithWindowScene() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) SetSelectedAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextFormattingCoordinator) ToggleFontPanel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


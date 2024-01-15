//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UITextInteraction : objc.NSObject
*/

type UITextInteraction struct {
  *objc.NSObject

}

func (r *UITextInteraction) TextInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInteraction) SetTextInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInteraction) TextInteractionMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInteraction) GesturesForFailureRequirements() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInteraction) TextInteractionForMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UITextInteraction) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


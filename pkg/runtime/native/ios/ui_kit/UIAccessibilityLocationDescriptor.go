//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIAccessibilityLocationDescriptor : objc.NSObject
*/

type UIAccessibilityLocationDescriptor struct {
  *objc.NSObject

}

func (r *UIAccessibilityLocationDescriptor) AttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) InitWithAttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) View() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) Point() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityLocationDescriptor) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


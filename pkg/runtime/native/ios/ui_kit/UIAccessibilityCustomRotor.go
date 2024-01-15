//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface UIKit.UIAccessibilityCustomRotor : objc.NSObject
*/

type UIAccessibilityCustomRotor struct {
  *objc.NSObject

}

func (r *UIAccessibilityCustomRotor) ItemSearchBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) SystemRotorType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) InitWithAttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) AttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) SetAttributedName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) InitWithSystemType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) SetName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotor) SetItemSearchBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


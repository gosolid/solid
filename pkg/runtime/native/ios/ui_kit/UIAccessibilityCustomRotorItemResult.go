//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIAccessibilityCustomRotorItemResult : objc.NSObject
*/

type UIAccessibilityCustomRotorItemResult struct {
  *objc.NSObject

}

func (r *UIAccessibilityCustomRotorItemResult) InitWithTargetElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorItemResult) TargetElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorItemResult) SetTargetElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorItemResult) TargetRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorItemResult) SetTargetRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


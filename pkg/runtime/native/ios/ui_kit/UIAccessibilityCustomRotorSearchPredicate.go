//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIAccessibilityCustomRotorSearchPredicate : objc.NSObject
*/

type UIAccessibilityCustomRotorSearchPredicate struct {
  *objc.NSObject

}

func (r *UIAccessibilityCustomRotorSearchPredicate) CurrentItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorSearchPredicate) SetCurrentItem() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorSearchPredicate) SearchDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityCustomRotorSearchPredicate) SetSearchDirection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


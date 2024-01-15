//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface AppKit.NSAccessibilityCustomRotorItemResult : objc.NSObject
*/

type NSAccessibilityCustomRotorItemResult struct {
  *objc.NSObject

}

func (r *NSAccessibilityCustomRotorItemResult) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) InitWithTargetElement() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) InitWithItemLoadingToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) TargetRange() (foundation.NSRange, error) {
  return foundation.NSRange{}, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) CustomLabel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) SetCustomLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) TargetElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) ItemLoadingToken() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorItemResult) SetTargetRange() error {
  return fmt.Errorf("unimplemented")
}


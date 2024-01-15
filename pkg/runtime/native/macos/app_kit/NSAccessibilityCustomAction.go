//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSAccessibilityCustomAction : objc.NSObject
*/

type NSAccessibilityCustomAction struct {
  *objc.NSObject

}

func (r *NSAccessibilityCustomAction) SetHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) Target() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) Selector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) SetSelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) Handler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomAction) SetTarget() error {
  return fmt.Errorf("unimplemented")
}


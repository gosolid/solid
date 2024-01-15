//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSAccessibilityCustomRotor : objc.NSObject
*/

type NSAccessibilityCustomRotor struct {
  *objc.NSObject

}

func (r *NSAccessibilityCustomRotor) InitWithRotorType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) SetType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) Label() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) SetLabel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) InitWithLabel() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) Type() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) ItemSearchDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) SetItemSearchDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) ItemLoadingDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotor) SetItemLoadingDelegate() error {
  return fmt.Errorf("unimplemented")
}


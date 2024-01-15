//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSAccessibilityCustomRotorSearchParameters : objc.NSObject
*/

type NSAccessibilityCustomRotorSearchParameters struct {
  *objc.NSObject

}

func (r *NSAccessibilityCustomRotorSearchParameters) CurrentItem() (*NSAccessibilityCustomRotorItemResult, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorSearchParameters) SetCurrentItem() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorSearchParameters) SearchDirection() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorSearchParameters) SetSearchDirection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorSearchParameters) FilterString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAccessibilityCustomRotorSearchParameters) SetFilterString() error {
  return fmt.Errorf("unimplemented")
}


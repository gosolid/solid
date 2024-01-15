//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIContentUnavailableImageProperties : objc.NSObject
*/

type UIContentUnavailableImageProperties struct {
  *objc.NSObject

}

func (r *UIContentUnavailableImageProperties) SetMaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) AccessibilityIgnoresInvertColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) PreferredSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) TintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) SetTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) MaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) SetPreferredSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) CornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) SetCornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIContentUnavailableImageProperties) SetAccessibilityIgnoresInvertColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


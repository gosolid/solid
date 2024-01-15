//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIListContentImageProperties : objc.NSObject
*/

type UIListContentImageProperties struct {
  *objc.NSObject

}

func (r *UIListContentImageProperties) SetAccessibilityIgnoresInvertColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) SetPreferredSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) SetTintColorTransformer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) CornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) SetMaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) ReservedLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) AccessibilityIgnoresInvertColors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) ResolvedTintColorForTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) PreferredSymbolConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) TintColorTransformer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) TintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) MaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) SetTintColor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) SetCornerRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIListContentImageProperties) SetReservedLayoutSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


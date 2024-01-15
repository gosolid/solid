//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIAccessibilityElement : UIKit.UIResponder
*/

type UIAccessibilityElement struct {
  *UIResponder

}

func (r *UIAccessibilityElement) SetAccessibilityLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityHint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetAccessibilityValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetAccessibilityTraits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetIsAccessibilityElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetAccessibilityFrameInContainerSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetAccessibilityContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetAccessibilityHint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) IsAccessibilityElement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) SetAccessibilityFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityTraits() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) AccessibilityFrameInContainerSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccessibilityElement) InitWithAccessibilityContainer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSLayoutGuide : objc.NSObject
*/

type NSLayoutGuide struct {
  *objc.NSObject
  *foundation.NSCoding
  *NSUserInterfaceItemIdentification
}

func (r *NSLayoutGuide) LeftAnchor() (*NSLayoutXAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) BottomAnchor() (*NSLayoutYAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) HeightAnchor() (*NSLayoutDimension, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) ConstraintsAffectingLayoutForOrientation() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) LeadingAnchor() (*NSLayoutXAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) WidthAnchor() (*NSLayoutDimension, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) CenterXAnchor() (*NSLayoutXAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) CenterYAnchor() (*NSLayoutYAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) Frame() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) RightAnchor() (*NSLayoutXAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) TopAnchor() (*NSLayoutYAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) HasAmbiguousLayout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) OwningView() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) SetOwningView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) TrailingAnchor() (*NSLayoutXAxisAnchor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) Identifier() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLayoutGuide) SetIdentifier() error {
  return fmt.Errorf("unimplemented")
}


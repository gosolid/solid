//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSPathControl : AppKit.NSControl
*/

type NSPathControl struct {
  *NSControl

}

func (r *NSPathControl) SetAllowedTypes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetPlaceholderAttributedString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetPathStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) ClickedPathItem() (*NSPathControlItem, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) Menu() (*NSMenu, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) AllowedTypes() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) PlaceholderAttributedString() (*foundation.NSAttributedString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetPathItems() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetBackgroundColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetMenu() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) PathStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) PathItems() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetDraggingSourceOperationMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) PlaceholderString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetPlaceholderString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) URL() (*foundation.NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) BackgroundColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPathControl) DoubleAction() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPathControl) SetDoubleAction() error {
  return fmt.Errorf("unimplemented")
}


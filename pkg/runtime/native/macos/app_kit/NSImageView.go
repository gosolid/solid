//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSImageView : AppKit.NSControl
*/

type NSImageView struct {
  *NSControl
  *NSAccessibilityImage
  *NSMenuItemValidation
}

func (r *NSImageView) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) ImageFrameStyle() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetSymbolConfiguration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetAnimates() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) IsEditable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SymbolConfiguration() (*NSImageSymbolConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageView) ContentTintColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageView) ImageDynamicRange() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageView) Animates() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetAllowsCutCopyPaste() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageView) ImageAlignment() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetImageAlignment() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) ImageScaling() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetImageFrameStyle() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetContentTintColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) DefaultPreferredImageDynamicRange() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetPreferredImageDynamicRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) ImageViewWithImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetEditable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetImageScaling() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) AllowsCutCopyPaste() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSImageView) SetDefaultPreferredImageDynamicRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSImageView) PreferredImageDynamicRange() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


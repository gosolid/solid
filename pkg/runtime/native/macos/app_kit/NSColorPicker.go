//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSColorPicker : objc.NSObject
*/

type NSColorPicker struct {
  *objc.NSObject
  *NSColorPickingDefault
}

func (r *NSColorPicker) InsertNewButtonImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) DetachColorList() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) ColorPanel() (*NSColorPanel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) ProvideNewButtonImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) ButtonToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) MinContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) InitWithPickerMask() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) AttachColorList() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) SetMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPicker) ViewSizeChanged() error {
  return fmt.Errorf("unimplemented")
}


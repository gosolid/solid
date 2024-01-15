//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol AppKit.NSColorPickingDefault
*/

type NSColorPickingDefault struct {

}

func (r *NSColorPickingDefault) SetMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) ButtonToolTip() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) MinContentSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) InitWithPickerMask() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) ProvideNewButtonImage() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) InsertNewButtonImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) ViewSizeChanged() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) AttachColorList() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) AlphaControlAddedOrRemoved() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSColorPickingDefault) DetachColorList() error {
  return fmt.Errorf("unimplemented")
}


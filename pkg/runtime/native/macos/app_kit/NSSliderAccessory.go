//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSSliderAccessory : objc.NSObject
*/

type NSSliderAccessory struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSSliderAccessory) AccessoryWithImage() (*NSSliderAccessory, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessory) Behavior() (*NSSliderAccessoryBehavior, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessory) SetBehavior() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessory) IsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessory) SetEnabled() error {
  return fmt.Errorf("unimplemented")
}


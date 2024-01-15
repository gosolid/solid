//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSSliderAccessoryBehavior : objc.NSObject
*/

type NSSliderAccessoryBehavior struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
}

func (r *NSSliderAccessoryBehavior) BehaviorWithHandler() (*NSSliderAccessoryBehavior, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessoryBehavior) HandleAction() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessoryBehavior) AutomaticBehavior() (*NSSliderAccessoryBehavior, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessoryBehavior) ValueStepBehavior() (*NSSliderAccessoryBehavior, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessoryBehavior) ValueResetBehavior() (*NSSliderAccessoryBehavior, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSliderAccessoryBehavior) BehaviorWithTarget() (*NSSliderAccessoryBehavior, error) {
  return nil, fmt.Errorf("unimplemented")
}


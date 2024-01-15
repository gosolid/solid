//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerAccessory : objc.NSObject
*/

type UIPointerAccessory struct {
  *objc.NSObject

}

func (r *UIPointerAccessory) ArrowAccessoryWithPosition() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) Shape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) Position() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) OrientationMatchesAngle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) SetOrientationMatchesAngle() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerAccessory) AccessoryWithShape() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


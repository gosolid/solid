//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerRegion : objc.NSObject
*/

type UIPointerRegion struct {
  *objc.NSObject

}

func (r *UIPointerRegion) RegionWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegion) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegion) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegion) Rect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegion) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegion) LatchingAxes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegion) SetLatchingAxes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


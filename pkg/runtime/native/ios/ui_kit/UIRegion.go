//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIRegion : objc.NSObject
*/

type UIRegion struct {
  *objc.NSObject

}

func (r *UIRegion) ContainsPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) InfiniteRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) InitWithRadius() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) InitWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) InverseRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) RegionByUnionWithRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) RegionByDifferenceFromRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIRegion) RegionByIntersectionWithRegion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


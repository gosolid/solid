//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIImageAsset : objc.NSObject
*/

type UIImageAsset struct {
  *objc.NSObject

}

func (r *UIImageAsset) UnregisterImageWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageAsset) ImageWithTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageAsset) UnregisterImageWithTraitCollection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageAsset) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageAsset) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageAsset) ImageWithConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIImageAsset) RegisterImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


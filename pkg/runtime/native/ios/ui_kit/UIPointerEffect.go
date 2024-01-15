//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerEffect : objc.NSObject
*/

type UIPointerEffect struct {
  *objc.NSObject

}

func (r *UIPointerEffect) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerEffect) Preview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerEffect) EffectWithPreview() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerEffect) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


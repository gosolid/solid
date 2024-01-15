//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerRegionRequest : objc.NSObject
*/

type UIPointerRegionRequest struct {
  *objc.NSObject

}

func (r *UIPointerRegionRequest) Location() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerRegionRequest) Modifiers() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


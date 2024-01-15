//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIAcceleration : objc.NSObject
*/

type UIAcceleration struct {
  *objc.NSObject

}

func (r *UIAcceleration) Z() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAcceleration) Timestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAcceleration) X() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAcceleration) Y() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


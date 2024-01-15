//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIAccelerometer : objc.NSObject
*/

type UIAccelerometer struct {
  *objc.NSObject

}

func (r *UIAccelerometer) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccelerometer) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccelerometer) SharedAccelerometer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccelerometer) UpdateInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIAccelerometer) SetUpdateInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


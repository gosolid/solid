//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIPointerLockState : objc.NSObject
*/

type UIPointerLockState struct {
  *objc.NSObject

}

func (r *UIPointerLockState) IsLocked() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerLockState) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIPointerLockState) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


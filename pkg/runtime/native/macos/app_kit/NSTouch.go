//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSTouch : objc.NSObject
*/

type NSTouch struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *NSTouch) Identity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouch) Phase() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTouch) NormalizedPosition() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSTouch) IsResting() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTouch) Device() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTouch) DeviceSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}


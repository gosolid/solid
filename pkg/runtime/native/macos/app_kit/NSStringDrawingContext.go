//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface AppKit.NSStringDrawingContext : objc.NSObject
*/

type NSStringDrawingContext struct {
  *objc.NSObject

}

func (r *NSStringDrawingContext) MinimumScaleFactor() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStringDrawingContext) SetMinimumScaleFactor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStringDrawingContext) ActualScaleFactor() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStringDrawingContext) TotalBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}


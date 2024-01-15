//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSTrackingArea : objc.NSObject
*/

type NSTrackingArea struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSTrackingArea) Rect() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSTrackingArea) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTrackingArea) Owner() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTrackingArea) UserInfo() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTrackingArea) InitWithRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


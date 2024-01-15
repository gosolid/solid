//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSShadow : objc.NSObject
*/

type NSShadow struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSShadow) Set() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSShadow) ShadowOffset() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSShadow) SetShadowOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSShadow) ShadowBlurRadius() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSShadow) SetShadowBlurRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSShadow) ShadowColor() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSShadow) SetShadowColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSShadow) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


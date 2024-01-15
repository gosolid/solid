//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSGradient : objc.NSObject
*/

type NSGradient struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *NSGradient) InitWithStartingColor() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGradient) InitWithColorsAndLocations() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGradient) DrawFromPoint() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGradient) DrawInRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGradient) DrawInBezierPath() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGradient) InterpolatedColorAtLocation() (*NSColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGradient) NumberOfColorStops() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSGradient) InitWithColors() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGradient) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGradient) DrawFromCenter() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGradient) GetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGradient) ColorSpace() (*NSColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}


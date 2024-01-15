//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSRulerMarker : objc.NSObject
*/

type NSRulerMarker struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSRulerMarker) SetMovable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) SetRemovable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) IsDragging() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) InitWithRulerView() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) SetMarkerLocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) Image() (*NSImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) ImageOrigin() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) SetImageOrigin() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) ThicknessRequiredInRuler() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) SetRepresentedObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) Ruler() (*NSRulerView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) MarkerLocation() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) IsMovable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) DrawRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) TrackMouse() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) IsRemovable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) ImageRectInRuler() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRulerMarker) RepresentedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


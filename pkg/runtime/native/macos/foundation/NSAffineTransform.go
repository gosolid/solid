//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAffineTransform : objc.NSObject
*/

type NSAffineTransform struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSAffineTransform) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) RotateByDegrees() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) ScaleXBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) PrependTransform() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) TransformStruct() (NSAffineTransformStruct, error) {
  return NSAffineTransformStruct{}, fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) Transform() (*NSAffineTransform, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) InitWithTransform() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) SetTransformStruct() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) TranslateXBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) RotateByRadians() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) TransformSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) ScaleBy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) TransformPoint() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) Invert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAffineTransform) AppendTransform() error {
  return fmt.Errorf("unimplemented")
}


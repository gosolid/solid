//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreImage.CIVector : objc.NSObject
*/

type CIVector struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSSecureCoding
}

func (r *CIVector) VectorWithX() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithCGAffineTransform() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithValues() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) CGAffineTransformValue() (core_foundation.CGAffineTransform, error) {
  return core_foundation.CGAffineTransform{}, fmt.Errorf("unimplemented")
}

func (r *CIVector) CGPointValue() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIVector) CGRectValue() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIVector) StringRepresentation() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithCGRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithCGRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) Y() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVector) W() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVector) Z() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithCGPoint() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithX() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithCGAffineTransform() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) ValueAtIndex() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVector) Count() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithValues() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithCGPoint() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) X() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


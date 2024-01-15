//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIVector : objc.NSObject
*/

type CIVector struct {
  *objc.NSObject

}

func (r *CIVector) VectorWithCGRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithCGPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) ValueAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) X() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) Y() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) CGPointValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) CGAffineTransformValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithCGAffineTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) CGRectValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) StringRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithValues() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithX() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithX() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithCGRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithCGAffineTransform() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) Z() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) W() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) VectorWithCGPoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIVector) InitWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


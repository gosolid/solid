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
interface CoreImage.CIFilterShape : objc.NSObject
*/

type CIFilterShape struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *CIFilterShape) TransformBy() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) InsetByX() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) UnionWith() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) UnionWithRect() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) ShapeWithRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) InitWithRect() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) IntersectWith() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) IntersectWithRect() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}


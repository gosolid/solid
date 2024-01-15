//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIFilterShape : objc.NSObject
*/

type CIFilterShape struct {
  *objc.NSObject

}

func (r *CIFilterShape) IntersectWith() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) InsetByX() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) UnionWith() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) UnionWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) IntersectWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) Extent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) ShapeWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) InitWithRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFilterShape) TransformBy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


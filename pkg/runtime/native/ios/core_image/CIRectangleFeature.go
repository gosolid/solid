//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIRectangleFeature : CoreImage.CIFeature
*/

type CIRectangleFeature struct {
  *CIFeature

}

func (r *CIRectangleFeature) BottomRight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRectangleFeature) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRectangleFeature) TopLeft() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRectangleFeature) TopRight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRectangleFeature) BottomLeft() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


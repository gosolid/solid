//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CITextFeature : CoreImage.CIFeature
*/

type CITextFeature struct {
  *CIFeature

}

func (r *CITextFeature) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextFeature) TopLeft() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextFeature) TopRight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextFeature) BottomLeft() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextFeature) BottomRight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITextFeature) SubFeatures() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


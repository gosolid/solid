//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIQRCodeFeature : CoreImage.CIFeature
*/

type CIQRCodeFeature struct {
  *CIFeature

}

func (r *CIQRCodeFeature) SymbolDescriptor() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) TopLeft() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) TopRight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) BottomLeft() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) BottomRight() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) MessageString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


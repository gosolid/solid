//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreImage.CIQRCodeFeature : CoreImage.CIFeature
*/

type CIQRCodeFeature struct {
  *CIFeature
  *foundation.NSSecureCoding
  *foundation.NSCopying
}

func (r *CIQRCodeFeature) MessageString() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) SymbolDescriptor() (*CIQRCodeDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) TopLeft() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) TopRight() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) BottomLeft() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeFeature) BottomRight() (core_foundation.CGPoint, error) {
  return core_foundation.CGPoint{}, fmt.Errorf("unimplemented")
}


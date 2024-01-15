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
interface CoreImage.CISampler : objc.NSObject
*/

type CISampler struct {
  *objc.NSObject
  *foundation.NSCopying
}

func (r *CISampler) SamplerWithImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISampler) InitWithImage() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISampler) Definition() (*CIFilterShape, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISampler) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreImage.CIImageAccumulator : objc.NSObject
*/

type CIImageAccumulator struct {
  *objc.NSObject

}

func (r *CIImageAccumulator) ImageAccumulatorWithExtent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) InitWithExtent() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Image() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) SetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Clear() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIImageAccumulator) Format() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


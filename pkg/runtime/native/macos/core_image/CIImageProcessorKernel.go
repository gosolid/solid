//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreImage.CIImageProcessorKernel : objc.NSObject
*/

type CIImageProcessorKernel struct {
  *objc.NSObject

}

func (r *CIImageProcessorKernel) OutputIsOpaque() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) SynchronizeInputs() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) RoiForInput() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) RoiTileArrayForInput() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) FormatForInputAtIndex() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) ApplyWithExtent() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) OutputFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


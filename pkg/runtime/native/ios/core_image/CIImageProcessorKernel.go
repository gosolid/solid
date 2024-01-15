//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIImageProcessorKernel : objc.NSObject
*/

type CIImageProcessorKernel struct {
  *objc.NSObject

}

func (r *CIImageProcessorKernel) FormatForInputAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) ApplyWithExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) OutputFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) OutputIsOpaque() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) SynchronizeInputs() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) RoiForInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIImageProcessorKernel) RoiTileArrayForInput() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


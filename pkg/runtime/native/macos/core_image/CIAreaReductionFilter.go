//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
protocol CoreImage.CIAreaReductionFilter
*/

type CIAreaReductionFilter struct {

}

func (r *CIAreaReductionFilter) SetExtent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaReductionFilter) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIAreaReductionFilter) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIAreaReductionFilter) Extent() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}


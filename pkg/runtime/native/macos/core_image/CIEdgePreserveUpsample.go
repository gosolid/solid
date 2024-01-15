//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIEdgePreserveUpsample
*/

type CIEdgePreserveUpsample struct {

}

func (r *CIEdgePreserveUpsample) SetLumaSigma() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) SmallImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) SetSmallImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) SpatialSigma() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) SetSpatialSigma() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdgePreserveUpsample) LumaSigma() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


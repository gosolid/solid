//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIConvolution
*/

type CIConvolution struct {

}

func (r *CIConvolution) Weights() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIConvolution) SetWeights() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIConvolution) Bias() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIConvolution) SetBias() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIConvolution) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIConvolution) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorMatrix
*/

type CIColorMatrix struct {

}

func (r *CIColorMatrix) BiasVector() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) RVector() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) GVector() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) BVector() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) SetBVector() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) AVector() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) SetRVector() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) SetGVector() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) SetAVector() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorMatrix) SetBiasVector() error {
  return fmt.Errorf("unimplemented")
}


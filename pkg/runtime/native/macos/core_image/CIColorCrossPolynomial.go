//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorCrossPolynomial
*/

type CIColorCrossPolynomial struct {

}

func (r *CIColorCrossPolynomial) SetGreenCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) BlueCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) SetBlueCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) RedCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) SetRedCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorCrossPolynomial) GreenCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}


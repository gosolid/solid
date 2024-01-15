//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorPolynomial
*/

type CIColorPolynomial struct {

}

func (r *CIColorPolynomial) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) SetRedCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) GreenCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) SetBlueCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) SetAlphaCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) RedCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) SetGreenCoefficients() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) BlueCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorPolynomial) AlphaCoefficients() (*CIVector, error) {
  return nil, fmt.Errorf("unimplemented")
}


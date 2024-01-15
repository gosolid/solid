//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIFalseColor
*/

type CIFalseColor struct {

}

func (r *CIFalseColor) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFalseColor) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFalseColor) Color0() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFalseColor) SetColor0() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIFalseColor) Color1() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIFalseColor) SetColor1() error {
  return fmt.Errorf("unimplemented")
}


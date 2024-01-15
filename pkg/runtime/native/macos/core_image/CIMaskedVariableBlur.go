//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMaskedVariableBlur
*/

type CIMaskedVariableBlur struct {

}

func (r *CIMaskedVariableBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMaskedVariableBlur) Mask() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMaskedVariableBlur) SetMask() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMaskedVariableBlur) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIMaskedVariableBlur) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIMaskedVariableBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}


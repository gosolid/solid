//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBoxBlur
*/

type CIBoxBlur struct {

}

func (r *CIBoxBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBoxBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBoxBlur) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBoxBlur) SetRadius() error {
  return fmt.Errorf("unimplemented")
}


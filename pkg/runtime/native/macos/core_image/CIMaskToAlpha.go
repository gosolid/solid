//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIMaskToAlpha
*/

type CIMaskToAlpha struct {

}

func (r *CIMaskToAlpha) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIMaskToAlpha) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


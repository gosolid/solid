//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIWhitePointAdjust
*/

type CIWhitePointAdjust struct {

}

func (r *CIWhitePointAdjust) SetColor() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIWhitePointAdjust) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIWhitePointAdjust) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIWhitePointAdjust) Color() (*CIColor, error) {
  return nil, fmt.Errorf("unimplemented")
}


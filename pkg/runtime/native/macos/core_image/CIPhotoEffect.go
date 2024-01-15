//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIPhotoEffect
*/

type CIPhotoEffect struct {

}

func (r *CIPhotoEffect) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPhotoEffect) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIPhotoEffect) Extrapolate() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPhotoEffect) SetExtrapolate() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIColorPosterize
*/

type CIColorPosterize struct {

}

func (r *CIColorPosterize) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIColorPosterize) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIColorPosterize) Levels() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIColorPosterize) SetLevels() error {
  return fmt.Errorf("unimplemented")
}


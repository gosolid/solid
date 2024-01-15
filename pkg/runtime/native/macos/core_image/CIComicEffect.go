//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIComicEffect
*/

type CIComicEffect struct {

}

func (r *CIComicEffect) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIComicEffect) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}


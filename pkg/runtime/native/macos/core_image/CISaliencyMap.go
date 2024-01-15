//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CISaliencyMap
*/

type CISaliencyMap struct {

}

func (r *CISaliencyMap) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CISaliencyMap) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}


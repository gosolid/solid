//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIEdges
*/

type CIEdges struct {

}

func (r *CIEdges) SetIntensity() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdges) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIEdges) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdges) Intensity() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIEdgeWork
*/

type CIEdgeWork struct {

}

func (r *CIEdgeWork) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdgeWork) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIEdgeWork) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIEdgeWork) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


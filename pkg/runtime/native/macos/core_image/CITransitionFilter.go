//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CITransitionFilter
*/

type CITransitionFilter struct {

}

func (r *CITransitionFilter) SetTargetImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITransitionFilter) Time() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CITransitionFilter) SetTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITransitionFilter) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CITransitionFilter) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CITransitionFilter) TargetImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}


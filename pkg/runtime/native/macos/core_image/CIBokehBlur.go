//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBokehBlur
*/

type CIBokehBlur struct {

}

func (r *CIBokehBlur) SetRadius() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) SetRingAmount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) SetRingSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) RingSize() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) Softness() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) SetSoftness() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) Radius() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIBokehBlur) RingAmount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


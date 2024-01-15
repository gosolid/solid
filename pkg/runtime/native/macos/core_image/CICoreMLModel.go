//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CICoreMLModel
*/

type CICoreMLModel struct {

}

func (r *CICoreMLModel) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) Model() (*MLModel, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) SetModel() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) HeadIndex() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) SetHeadIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) SoftmaxNormalization() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICoreMLModel) SetSoftmaxNormalization() error {
  return fmt.Errorf("unimplemented")
}


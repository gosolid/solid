//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIDocumentEnhancer
*/

type CIDocumentEnhancer struct {

}

func (r *CIDocumentEnhancer) InputImage() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDocumentEnhancer) SetInputImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIDocumentEnhancer) Amount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDocumentEnhancer) SetAmount() error {
  return fmt.Errorf("unimplemented")
}


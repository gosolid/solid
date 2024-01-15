//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol CoreImage.CIBarcodeGenerator
*/

type CIBarcodeGenerator struct {

}

func (r *CIBarcodeGenerator) BarcodeDescriptor() (*CIBarcodeDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBarcodeGenerator) SetBarcodeDescriptor() error {
  return fmt.Errorf("unimplemented")
}


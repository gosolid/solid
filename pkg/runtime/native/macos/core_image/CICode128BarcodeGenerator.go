//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol CoreImage.CICode128BarcodeGenerator
*/

type CICode128BarcodeGenerator struct {

}

func (r *CICode128BarcodeGenerator) SetBarcodeHeight() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICode128BarcodeGenerator) Message() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CICode128BarcodeGenerator) SetMessage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICode128BarcodeGenerator) QuietSpace() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CICode128BarcodeGenerator) SetQuietSpace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CICode128BarcodeGenerator) BarcodeHeight() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}


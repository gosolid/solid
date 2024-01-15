//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
protocol CoreImage.CIQRCodeGenerator
*/

type CIQRCodeGenerator struct {

}

func (r *CIQRCodeGenerator) Message() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeGenerator) SetMessage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIQRCodeGenerator) CorrectionLevel() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeGenerator) SetCorrectionLevel() error {
  return fmt.Errorf("unimplemented")
}


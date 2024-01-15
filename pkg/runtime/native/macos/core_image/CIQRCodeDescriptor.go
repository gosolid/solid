//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreImage.CIQRCodeDescriptor : CoreImage.CIBarcodeDescriptor
*/

type CIQRCodeDescriptor struct {
  *CIBarcodeDescriptor

}

func (r *CIQRCodeDescriptor) InitWithPayload() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) DescriptorWithPayload() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) ErrorCorrectedPayload() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) SymbolVersion() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) MaskPattern() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) ErrorCorrectionLevel() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


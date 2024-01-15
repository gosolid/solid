//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIQRCodeDescriptor : CoreImage.CIBarcodeDescriptor
*/

type CIQRCodeDescriptor struct {
  *CIBarcodeDescriptor

}

func (r *CIQRCodeDescriptor) ErrorCorrectedPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) SymbolVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) MaskPattern() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) ErrorCorrectionLevel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) InitWithPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIQRCodeDescriptor) DescriptorWithPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


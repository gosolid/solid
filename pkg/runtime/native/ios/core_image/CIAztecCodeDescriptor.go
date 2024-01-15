//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIAztecCodeDescriptor : CoreImage.CIBarcodeDescriptor
*/

type CIAztecCodeDescriptor struct {
  *CIBarcodeDescriptor

}

func (r *CIAztecCodeDescriptor) InitWithPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIAztecCodeDescriptor) DescriptorWithPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


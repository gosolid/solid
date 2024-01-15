//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIPDF417CodeDescriptor : CoreImage.CIBarcodeDescriptor
*/

type CIPDF417CodeDescriptor struct {
  *CIBarcodeDescriptor

}

func (r *CIPDF417CodeDescriptor) InitWithPayload() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIPDF417CodeDescriptor) DescriptorWithPayload() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


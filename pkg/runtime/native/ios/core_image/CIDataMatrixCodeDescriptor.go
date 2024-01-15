//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIDataMatrixCodeDescriptor : CoreImage.CIBarcodeDescriptor
*/

type CIDataMatrixCodeDescriptor struct {
  *CIBarcodeDescriptor

}

func (r *CIDataMatrixCodeDescriptor) InitWithPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) DescriptorWithPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) ErrorCorrectedPayload() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) RowCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) ColumnCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) EccVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


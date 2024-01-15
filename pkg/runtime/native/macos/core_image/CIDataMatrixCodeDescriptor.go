//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreImage.CIDataMatrixCodeDescriptor : CoreImage.CIBarcodeDescriptor
*/

type CIDataMatrixCodeDescriptor struct {
  *CIBarcodeDescriptor

}

func (r *CIDataMatrixCodeDescriptor) EccVersion() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) InitWithPayload() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) DescriptorWithPayload() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) ErrorCorrectedPayload() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) RowCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIDataMatrixCodeDescriptor) ColumnCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}


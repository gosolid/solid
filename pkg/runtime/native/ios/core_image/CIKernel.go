//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIKernel : objc.NSObject
*/

type CIKernel struct {
  *objc.NSObject

}

func (r *CIKernel) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelsWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelsWithMetalString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelWithString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelWithFunctionName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelNamesFromMetalLibraryData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) SetROISelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) ApplyWithExtent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


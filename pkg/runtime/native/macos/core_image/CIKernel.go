//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface CoreImage.CIKernel : objc.NSObject
*/

type CIKernel struct {
  *objc.NSObject

}

func (r *CIKernel) SetROISelector() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIKernel) ApplyWithExtent() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelsWithString() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelsWithMetalString() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelWithFunctionName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIKernel) KernelNamesFromMetalLibraryData() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIWarpKernel : CoreImage.CIKernel
*/

type CIWarpKernel struct {
  *CIKernel

}

func (r *CIWarpKernel) KernelWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIWarpKernel) ApplyWithExtent() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}


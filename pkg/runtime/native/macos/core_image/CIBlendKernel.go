//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreImage.CIBlendKernel : CoreImage.CIColorKernel
*/

type CIBlendKernel struct {
  *CIColorKernel

}

func (r *CIBlendKernel) KernelWithString() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIBlendKernel) ApplyWithForeground() (*CIImage, error) {
  return nil, fmt.Errorf("unimplemented")
}


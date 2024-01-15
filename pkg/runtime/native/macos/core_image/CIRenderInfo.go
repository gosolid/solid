//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface CoreImage.CIRenderInfo : objc.NSObject
*/

type CIRenderInfo struct {
  *objc.NSObject

}

func (r *CIRenderInfo) KernelExecutionTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRenderInfo) KernelCompileTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRenderInfo) PassCount() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIRenderInfo) PixelsProcessed() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}


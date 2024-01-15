//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIRenderInfo : objc.NSObject
*/

type CIRenderInfo struct {
  *objc.NSObject

}

func (r *CIRenderInfo) KernelExecutionTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderInfo) KernelCompileTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderInfo) PassCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIRenderInfo) PixelsProcessed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


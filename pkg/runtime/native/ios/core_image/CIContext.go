//js:package native/ios/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface CoreImage.CIContext : objc.NSObject
*/

type CIContext struct {
  *objc.NSObject

}

func (r *CIContext) InitWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithEAGLContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) WorkingFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithMTLDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) DrawImage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ClearCaches() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) OutputImageMaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithMTLCommandQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) CreateCGLayerWithSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ReclaimResources() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) WorkingColorSpace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithCGContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) Context() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) Render() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) InputImageMaximumSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


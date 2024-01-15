//js:package native/macos/core-image
package core_image

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface CoreImage.CIContext : objc.NSObject
*/

type CIContext struct {
  *objc.NSObject

}

func (r *CIContext) CreateCGLayerWithSize() (*core_graphics.CGLayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ReclaimResources() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIContext) WorkingFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithCGLContext() (*CIContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) Context() (*CIContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) InputImageMaximumSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CIContext) OutputImageMaximumSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithCGContext() (*CIContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithOptions() (*CIContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) InitWithOptions() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithMTLDevice() (*CIContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) DrawImage() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIContext) ContextWithMTLCommandQueue() (*CIContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CIContext) Render() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIContext) ClearCaches() error {
  return fmt.Errorf("unimplemented")
}

func (r *CIContext) WorkingColorSpace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}


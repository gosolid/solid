//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface QuartzCore.CARenderer : objc.NSObject
*/

type CARenderer struct {
  *objc.NSObject

}

func (r *CARenderer) SetDestination() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARenderer) Layer() (*CALayer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) SetBounds() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARenderer) RendererWithMTLTexture() (*CARenderer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) UpdateBounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CARenderer) NextFrameTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CARenderer) EndFrame() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARenderer) BeginFrameAtTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARenderer) AddUpdateRect() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARenderer) SetLayer() error {
  return fmt.Errorf("unimplemented")
}

func (r *CARenderer) Bounds() (core_foundation.CGRect, error) {
  return core_foundation.CGRect{}, fmt.Errorf("unimplemented")
}

func (r *CARenderer) RendererWithCGLContext() (*CARenderer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) Render() error {
  return fmt.Errorf("unimplemented")
}


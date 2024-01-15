//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/open_gl"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
)

/*
interface QuartzCore.CAOpenGLLayer : QuartzCore.CALayer
*/

type CAOpenGLLayer struct {
  *CALayer

}

func (r *CAOpenGLLayer) WantsExtendedDynamicRangeContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) DrawInCGLContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) CopyCGLPixelFormatForDisplayMask() (*open_gl.CGLPixelFormatObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) ReleaseCGLPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) ReleaseCGLContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) IsAsynchronous() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) Colorspace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) SetColorspace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) SetWantsExtendedDynamicRangeContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) CopyCGLContextForPixelFormat() (*open_gl.CGLContextObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAOpenGLLayer) SetAsynchronous() error {
  return fmt.Errorf("unimplemented")
}


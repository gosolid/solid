//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface AppKit.NSOpenGLView : AppKit.NSView
*/

type NSOpenGLView struct {
  *NSView

}

func (r *NSOpenGLView) DefaultPixelFormat() (*NSOpenGLPixelFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) InitWithFrame() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) OpenGLContext() (*NSOpenGLContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) Reshape() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) SetWantsBestResolutionOpenGLSurface() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) SetWantsExtendedDynamicRangeOpenGLSurface() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) Update() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) PixelFormat() (*NSOpenGLPixelFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) SetPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) WantsBestResolutionOpenGLSurface() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) ClearGLContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) PrepareOpenGL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) SetOpenGLContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLView) WantsExtendedDynamicRangeOpenGLSurface() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


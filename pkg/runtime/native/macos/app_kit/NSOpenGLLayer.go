//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/quartz_core"
  "fmt"
)

/*
interface AppKit.NSOpenGLLayer : QuartzCore.CAOpenGLLayer
*/

type NSOpenGLLayer struct {
  *quartz_core.CAOpenGLLayer

}

func (r *NSOpenGLLayer) CanDrawInOpenGLContext() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) SetOpenGLPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) SetOpenGLContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) OpenGLContext() (*NSOpenGLContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) OpenGLPixelFormatForDisplayMask() (*NSOpenGLPixelFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) OpenGLContextForPixelFormat() (*NSOpenGLContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) DrawInOpenGLContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLLayer) OpenGLPixelFormat() (*NSOpenGLPixelFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}


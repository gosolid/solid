//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/open_gl"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSOpenGLContext : objc.NSObject
*/

type NSOpenGLContext struct {
  *objc.NSObject
  *foundation.NSLocking
}

func (r *NSOpenGLContext) ClearDrawable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) MakeCurrentContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) GetValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) InitWithFormat() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) SetView() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) SetFullScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) Update() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) PixelFormat() (*NSOpenGLPixelFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) CurrentContext() (*NSOpenGLContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) SetCurrentVirtualScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) InitWithCGLContextObj() (*NSOpenGLContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) SetOffScreen() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) FlushBuffer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) ClearCurrentContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) CopyAttributesFromContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) View() (*NSView, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) SetValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) CreateTexture() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) CurrentVirtualScreen() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLContext) CGLContextObj() (*open_gl.CGLContextObject, error) {
  return nil, fmt.Errorf("unimplemented")
}


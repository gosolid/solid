//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/open_gl"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface AppKit.NSOpenGLPixelFormat : objc.NSObject
*/

type NSOpenGLPixelFormat struct {
  *objc.NSObject
  *foundation.NSCoding
}

func (r *NSOpenGLPixelFormat) CGLPixelFormatObj() (*open_gl.CGLPixelFormatObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) InitWithCGLPixelFormatObj() (*NSOpenGLPixelFormat, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) InitWithAttributes() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) Attributes() (*foundation.NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) SetAttributes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) GetValues() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelFormat) NumberOfVirtualScreens() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


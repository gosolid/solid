//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/open_gl"
)

/*
interface AppKit.NSOpenGLPixelBuffer : objc.NSObject
*/

type NSOpenGLPixelBuffer struct {
  *objc.NSObject

}

func (r *NSOpenGLPixelBuffer) TextureInternalFormat() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) TextureMaxMipMapLevel() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) InitWithTextureTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) InitWithCGLPBufferObj() (*NSOpenGLPixelBuffer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) CGLPBufferObj() (*open_gl.CGLPBufferObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) PixelsWide() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) PixelsHigh() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSOpenGLPixelBuffer) TextureTarget() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}


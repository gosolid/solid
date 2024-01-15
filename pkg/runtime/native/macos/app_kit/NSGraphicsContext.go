//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface AppKit.NSGraphicsContext : objc.NSObject
*/

type NSGraphicsContext struct {
  *objc.NSObject

}

func (r *NSGraphicsContext) GraphicsContextWithBitmapImageRep() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) CGContext() (*core_graphics.CGContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) GraphicsContextWithAttributes() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) SetCurrentContext() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) Attributes() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) CurrentContext() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) GraphicsContextWithCGContext() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) CurrentContextDrawingToScreen() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) IsFlipped() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) GraphicsContextWithWindow() (*NSGraphicsContext, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) RestoreGraphicsState() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) FlushGraphics() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) IsDrawingToScreen() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSGraphicsContext) SaveGraphicsState() error {
  return fmt.Errorf("unimplemented")
}


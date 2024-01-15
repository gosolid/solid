//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_graphics"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface QuartzCore.CAMetalLayer : QuartzCore.CALayer
*/

type CAMetalLayer struct {
  *CALayer

}

func (r *CAMetalLayer) NextDrawable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetPixelFormat() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDrawableSize() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetMaximumDrawableCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) WantsExtendedDynamicRangeContent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetAllowsNextDrawableTimeout() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) MaximumDrawableCount() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDeveloperHUDProperties() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) FramebufferOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) DrawableSize() (core_foundation.CGSize, error) {
  return core_foundation.CGSize{}, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetWantsExtendedDynamicRangeContent() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetFramebufferOnly() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetColorspace() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDisplaySyncEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) AllowsNextDrawableTimeout() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetPresentsWithTransaction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetEDRMetadata() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) PreferredDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) Colorspace() (*core_graphics.CGColorSpace, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) DisplaySyncEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) DeveloperHUDProperties() (*foundation.NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDevice() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) EDRMetadata() (*CAEDRMetadata, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) PixelFormat() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) PresentsWithTransaction() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface QuartzCore.CAMetalLayer : QuartzCore.CALayer
*/

type CAMetalLayer struct {
  *CALayer

}

func (r *CAMetalLayer) Device() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) MaximumDrawableCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetMaximumDrawableCount() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetColorspace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) NextDrawable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) PreferredDevice() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) PixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) PresentsWithTransaction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) AllowsNextDrawableTimeout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) DeveloperHUDProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDrawableSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) Colorspace() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) WantsExtendedDynamicRangeContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetEDRMetadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetAllowsNextDrawableTimeout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetPixelFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) FramebufferOnly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetFramebufferOnly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) DrawableSize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetPresentsWithTransaction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetWantsExtendedDynamicRangeContent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) EDRMetadata() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalLayer) SetDeveloperHUDProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


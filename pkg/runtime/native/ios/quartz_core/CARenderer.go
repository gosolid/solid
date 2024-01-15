//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CARenderer : objc.NSObject
*/

type CARenderer struct {
  *objc.NSObject

}

func (r *CARenderer) NextFrameTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) EndFrame() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) RendererWithMTLTexture() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) BeginFrameAtTime() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) UpdateBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) Layer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) SetLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) Bounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) SetBounds() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) AddUpdateRect() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) Render() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CARenderer) SetDestination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


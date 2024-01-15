//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CAMetalDisplayLink : objc.NSObject
*/

type CAMetalDisplayLink struct {
  *objc.NSObject

}

func (r *CAMetalDisplayLink) InitWithMetalLayer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) AddToRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) RemoveFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) PreferredFrameLatency() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetPreferredFrameLatency() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) PreferredFrameRateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetPreferredFrameRateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) IsPaused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetPaused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


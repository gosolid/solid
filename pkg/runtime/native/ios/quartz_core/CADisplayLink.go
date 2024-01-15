//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CADisplayLink : objc.NSObject
*/

type CADisplayLink struct {
  *objc.NSObject

}

func (r *CADisplayLink) DisplayLinkWithTarget() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) Duration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) IsPaused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetPreferredFramesPerSecond() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) TargetTimestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetFrameInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) AddToRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) RemoveFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetPaused() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) PreferredFrameRateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetPreferredFrameRateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) Timestamp() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) FrameInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) PreferredFramesPerSecond() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


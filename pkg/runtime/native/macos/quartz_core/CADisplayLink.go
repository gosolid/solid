//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface QuartzCore.CADisplayLink : objc.NSObject
*/

type CADisplayLink struct {
  *objc.NSObject

}

func (r *CADisplayLink) Duration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) IsPaused() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetPaused() error {
  return fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) PreferredFrameRateRange() (CAFrameRateRange, error) {
  return CAFrameRateRange{}, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) FrameInterval() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetPreferredFramesPerSecond() error {
  return fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) DisplayLinkWithTarget() (*CADisplayLink, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) AddToRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) Timestamp() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetFrameInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) PreferredFramesPerSecond() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) TargetTimestamp() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CADisplayLink) SetPreferredFrameRateRange() error {
  return fmt.Errorf("unimplemented")
}


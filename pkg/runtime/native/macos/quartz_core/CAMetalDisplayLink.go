//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface QuartzCore.CAMetalDisplayLink : objc.NSObject
*/

type CAMetalDisplayLink struct {
  *objc.NSObject

}

func (r *CAMetalDisplayLink) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) PreferredFrameLatency() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) PreferredFrameRateRange() (CAFrameRateRange, error) {
  return CAFrameRateRange{}, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetPaused() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) InitWithMetalLayer() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetPreferredFrameLatency() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) SetPreferredFrameRateRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) IsPaused() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAMetalDisplayLink) AddToRunLoop() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface QuartzCore.CAAnimation : objc.NSObject
*/

type CAAnimation struct {
  *objc.NSObject
  *foundation.NSSecureCoding
  *foundation.NSCopying
  *CAMediaTiming
  *CAAction
}

func (r *CAAnimation) Animation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetTimingFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetRemovedOnCompletion() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAAnimation) PreferredFrameRateRange() (CAFrameRateRange, error) {
  return CAFrameRateRange{}, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) DefaultValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) TimingFunction() (*CAMediaTimingFunction, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) IsRemovedOnCompletion() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetPreferredFrameRateRange() error {
  return fmt.Errorf("unimplemented")
}


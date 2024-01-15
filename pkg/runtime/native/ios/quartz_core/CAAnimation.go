//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CAAnimation : objc.NSObject
*/

type CAAnimation struct {
  *objc.NSObject

}

func (r *CAAnimation) ShouldArchiveValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) TimingFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) IsRemovedOnCompletion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) PreferredFrameRateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) Animation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) DefaultValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetPreferredFrameRateRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetTimingFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAAnimation) SetRemovedOnCompletion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


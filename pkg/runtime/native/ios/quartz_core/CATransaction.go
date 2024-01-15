//js:package native/ios/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface QuartzCore.CATransaction : objc.NSObject
*/

type CATransaction struct {
  *objc.NSObject

}

func (r *CATransaction) Begin() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) Flush() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) Lock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetDisableActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) Batch() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetAnimationDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetAnimationTimingFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) CompletionBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetCompletionBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) Commit() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) Unlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) AnimationDuration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) AnimationTimingFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) DisableActions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) ValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


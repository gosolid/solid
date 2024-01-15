//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface QuartzCore.CATransaction : objc.NSObject
*/

type CATransaction struct {
  *objc.NSObject

}

func (r *CATransaction) SetAnimationDuration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) Begin() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) Lock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) AnimationDuration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CATransaction) CompletionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetCompletionBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) Commit() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) Flush() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) AnimationTimingFunction() (*CAMediaTimingFunction, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetAnimationTimingFunction() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetDisableActions() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) SetValue() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) Unlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *CATransaction) ValueForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CATransaction) Batch() error {
  return fmt.Errorf("unimplemented")
}


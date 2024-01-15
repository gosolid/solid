//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface AppKit.NSAnimation : objc.NSObject
*/

type NSAnimation struct {
  *objc.NSObject
  *foundation.NSCopying
  *foundation.NSCoding
}

func (r *NSAnimation) Duration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) FrameRate() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) ClearStartAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) CurrentProgress() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetCurrentProgress() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetAnimationBlockingMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetFrameRate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) ProgressMarks() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) InitWithDuration() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) ClearStopAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) StartWhenAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) IsAnimating() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) AnimationBlockingMode() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) AddProgressMark() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) StartAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetDuration() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) AnimationCurve() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) RunLoopModesForAnimating() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) CurrentValue() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) RemoveProgressMark() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetAnimationCurve() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAnimation) SetProgressMarks() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) StopAnimation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSAnimation) StopWhenAnimation() error {
  return fmt.Errorf("unimplemented")
}


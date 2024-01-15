//js:package native/macos/quartz-core
package quartz_core

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
protocol QuartzCore.CAMediaTiming
*/

type CAMediaTiming struct {

}

func (r *CAMediaTiming) SetRepeatDuration() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) Autoreverses() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) RepeatCount() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) FillMode() (**foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) TimeOffset() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetRepeatCount() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetAutoreverses() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetFillMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) BeginTime() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetBeginTime() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) Speed() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetSpeed() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetTimeOffset() error {
  return fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) RepeatDuration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) Duration() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *CAMediaTiming) SetDuration() error {
  return fmt.Errorf("unimplemented")
}


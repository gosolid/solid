//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSTimer : objc.NSObject
*/

type NSTimer struct {
  *objc.NSObject

}

func (r *NSTimer) TimerWithTimeInterval() (*NSTimer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) InitWithFireDate() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) Fire() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTimer) FireDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) Tolerance() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTimer) SetTolerance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTimer) IsValid() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSTimer) ScheduledTimerWithTimeInterval() (*NSTimer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTimer) SetFireDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSTimer) TimeInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSTimer) UserInfo() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSBackgroundActivityScheduler : objc.NSObject
*/

type NSBackgroundActivityScheduler struct {
  *objc.NSObject

}

func (r *NSBackgroundActivityScheduler) InitWithIdentifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) Identifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) QualityOfService() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) ShouldDefer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) SetInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) SetQualityOfService() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) Interval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) SetTolerance() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) ScheduleWithBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) Repeats() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) SetRepeats() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSBackgroundActivityScheduler) Tolerance() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


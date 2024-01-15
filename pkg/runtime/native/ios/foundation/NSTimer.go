//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSTimer : objc.NSObject
*/

type NSTimer struct {
  *objc.NSObject

}

func (r *NSTimer) TimerWithTimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) Fire() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) Tolerance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) SetTolerance() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) ScheduledTimerWithTimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) InitWithFireDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) FireDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) SetFireDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) TimeInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSTimer) IsValid() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


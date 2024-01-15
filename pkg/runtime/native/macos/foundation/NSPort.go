//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPort : objc.NSObject
*/

type NSPort struct {
  *objc.NSObject
  *NSCopying
  *NSCoding
}

func (r *NSPort) Port() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPort) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPort) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) AddConnection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPort) RemoveConnection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPort) ReservedSpaceLength() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSPort) ScheduleInRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPort) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPort) SendBeforeDate() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPort) IsValid() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


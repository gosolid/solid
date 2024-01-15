//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSStream : objc.NSObject
*/

type NSStream struct {
  *objc.NSObject

}

func (r *NSStream) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStream) StreamStatus() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSStream) Open() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStream) PropertyForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) SetProperty() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSStream) ScheduleInRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStream) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) Close() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStream) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSStream) StreamError() (*NSError, error) {
  return nil, fmt.Errorf("unimplemented")
}


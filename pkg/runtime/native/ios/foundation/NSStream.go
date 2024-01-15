//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSStream : objc.NSObject
*/

type NSStream struct {
  *objc.NSObject

}

func (r *NSStream) StreamStatus() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) StreamError() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) ScheduleInRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) RemoveFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) SetProperty() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) Open() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) Close() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSStream) PropertyForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


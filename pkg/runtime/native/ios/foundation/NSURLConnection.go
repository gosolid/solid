//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLConnection : objc.NSObject
*/

type NSURLConnection struct {
  *objc.NSObject

}

func (r *NSURLConnection) InitWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) ConnectionWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) UnscheduleFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) CanHandleRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) CurrentRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) Start() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) Cancel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) ScheduleInRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) SetDelegateQueue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) OriginalRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


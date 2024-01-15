//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLConnection : objc.NSObject
*/

type NSURLConnection struct {
  *objc.NSObject

}

func (r *NSURLConnection) Start() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) Cancel() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) ScheduleInRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) SetDelegateQueue() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) OriginalRequest() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) InitWithRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) ConnectionWithRequest() (*NSURLConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) UnscheduleFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) CanHandleRequest() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLConnection) CurrentRequest() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}


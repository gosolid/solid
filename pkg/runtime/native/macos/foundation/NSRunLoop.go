//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
interface Foundation.NSRunLoop : objc.NSObject
*/

type NSRunLoop struct {
  *objc.NSObject

}

func (r *NSRunLoop) CurrentRunLoop() (*NSRunLoop, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) CurrentMode() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) GetCFRunLoop() (*core_foundation.CFRunLoop, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) AddPort() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) RemovePort() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) MainRunLoop() (*NSRunLoop, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) AddTimer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) LimitDateForMode() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) AcceptInputForMode() error {
  return fmt.Errorf("unimplemented")
}


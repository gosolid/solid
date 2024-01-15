//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSRunLoop : objc.NSObject
*/

type NSRunLoop struct {
  *objc.NSObject

}

func (r *NSRunLoop) RemovePort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) LimitDateForMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) AcceptInputForMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) CurrentMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) GetCFRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) AddTimer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) AddPort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) CurrentRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSRunLoop) MainRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


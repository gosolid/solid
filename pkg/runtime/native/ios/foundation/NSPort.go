//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPort : objc.NSObject
*/

type NSPort struct {
  *objc.NSObject

}

func (r *NSPort) Port() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) RemoveFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) IsValid() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) ScheduleInRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) SendBeforeDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPort) ReservedSpaceLength() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


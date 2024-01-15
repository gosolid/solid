//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMachPort : Foundation.NSPort
*/

type NSMachPort struct {
  *NSPort

}

func (r *NSMachPort) PortWithMachPort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) InitWithMachPort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) ScheduleInRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) RemoveFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) MachPort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


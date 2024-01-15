//js:package native/macos/foundation
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

func (r *NSMachPort) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMachPort) MachPort() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) PortWithMachPort() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) InitWithMachPort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMachPort) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachPort) ScheduleInRunLoop() error {
  return fmt.Errorf("unimplemented")
}


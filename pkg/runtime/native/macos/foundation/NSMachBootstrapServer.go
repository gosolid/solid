//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMachBootstrapServer : Foundation.NSPortNameServer
*/

type NSMachBootstrapServer struct {
  *NSPortNameServer

}

func (r *NSMachBootstrapServer) SharedInstance() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachBootstrapServer) PortForName() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMachBootstrapServer) RegisterPort() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMachBootstrapServer) ServicePortWithName() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}


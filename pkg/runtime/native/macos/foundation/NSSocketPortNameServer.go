//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSSocketPortNameServer : Foundation.NSPortNameServer
*/

type NSSocketPortNameServer struct {
  *NSPortNameServer

}

func (r *NSSocketPortNameServer) RemovePortForName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSSocketPortNameServer) DefaultNameServerPortNumber() (uint16, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSocketPortNameServer) SetDefaultNameServerPortNumber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSSocketPortNameServer) SharedInstance() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPortNameServer) PortForName() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPortNameServer) RegisterPort() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


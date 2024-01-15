//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSSocketPort : Foundation.NSPort
*/

type NSSocketPort struct {
  *NSPort

}

func (r *NSSocketPort) ProtocolFamily() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) SocketType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) Protocol() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) InitRemoteWithProtocolFamily() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) InitWithTCPPort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) InitWithProtocolFamily() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) InitRemoteWithTCPPort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) Address() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) Socket() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSSocketPort) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


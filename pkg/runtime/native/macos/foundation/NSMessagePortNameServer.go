//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMessagePortNameServer : Foundation.NSPortNameServer
*/

type NSMessagePortNameServer struct {
  *NSPortNameServer

}

func (r *NSMessagePortNameServer) SharedInstance() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMessagePortNameServer) PortForName() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}


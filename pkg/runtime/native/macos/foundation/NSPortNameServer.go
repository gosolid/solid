//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPortNameServer : objc.NSObject
*/

type NSPortNameServer struct {
  *objc.NSObject

}

func (r *NSPortNameServer) SystemDefaultPortNameServer() (*NSPortNameServer, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortNameServer) PortForName() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortNameServer) RegisterPort() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPortNameServer) RemovePortForName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


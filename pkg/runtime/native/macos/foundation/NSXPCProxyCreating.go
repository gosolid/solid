//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSXPCProxyCreating
*/

type NSXPCProxyCreating struct {

}

func (r *NSXPCProxyCreating) RemoteObjectProxy() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCProxyCreating) RemoteObjectProxyWithErrorHandler() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCProxyCreating) SynchronousRemoteObjectProxyWithErrorHandler() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


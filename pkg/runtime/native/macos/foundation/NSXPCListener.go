//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSXPCListener : objc.NSObject
*/

type NSXPCListener struct {
  *objc.NSObject

}

func (r *NSXPCListener) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Endpoint() (*NSXPCListenerEndpoint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) ServiceListener() (*NSXPCListener, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) AnonymousListener() (*NSXPCListener, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) InitWithMachServiceName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Resume() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Suspend() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Activate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) SetConnectionCodeSigningRequirement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


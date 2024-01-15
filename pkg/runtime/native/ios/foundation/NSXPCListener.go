//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSXPCListener : objc.NSObject
*/

type NSXPCListener struct {
  *objc.NSObject

}

func (r *NSXPCListener) Suspend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Endpoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) InitWithMachServiceName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Resume() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Activate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) SetConnectionCodeSigningRequirement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) ServiceListener() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCListener) AnonymousListener() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


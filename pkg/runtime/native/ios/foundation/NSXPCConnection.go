//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSXPCConnection : objc.NSObject
*/

type NSXPCConnection struct {
  *objc.NSObject

}

func (r *NSXPCConnection) InitWithMachServiceName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) AuditSessionIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ExportedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetExportedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetRemoteObjectInterface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) RemoteObjectProxy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SynchronousRemoteObjectProxyWithErrorHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Endpoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) CurrentConnection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) RemoteObjectInterface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InterruptionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InvalidationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InitWithListenerEndpoint() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetExportedInterface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) RemoteObjectProxyWithErrorHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ProcessIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetInterruptionHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Activate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetCodeSigningRequirement() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Suspend() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Invalidate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ScheduleSendBarrierBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ServiceName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ExportedInterface() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetInvalidationHandler() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InitWithServiceName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Resume() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) EffectiveUserIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) EffectiveGroupIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


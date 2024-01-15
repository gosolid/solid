//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSXPCConnection : objc.NSObject
*/

type NSXPCConnection struct {
  *objc.NSObject
  *NSXPCProxyCreating
}

func (r *NSXPCConnection) RemoteObjectProxy() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) AuditSessionIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetRemoteObjectInterface() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InterruptionHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InvalidationHandler() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InitWithMachServiceName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Suspend() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ServiceName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetExportedInterface() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ExportedObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Resume() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetInvalidationHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SynchronousRemoteObjectProxyWithErrorHandler() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) RemoteObjectInterface() (*NSXPCInterface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ProcessIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) CurrentConnection() (*NSXPCConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ScheduleSendBarrierBlock() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetExportedObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) EffectiveUserIdentifier() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) EffectiveGroupIdentifier() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Activate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetCodeSigningRequirement() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) Endpoint() (*NSXPCListenerEndpoint, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) ExportedInterface() (*NSXPCInterface, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) SetInterruptionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InitWithServiceName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) InitWithListenerEndpoint() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSXPCConnection) RemoteObjectProxyWithErrorHandler() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}


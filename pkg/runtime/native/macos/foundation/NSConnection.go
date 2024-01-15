//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSConnection : objc.NSObject
*/

type NSConnection struct {
  *objc.NSObject

}

func (r *NSConnection) ConnectionWithRegisteredName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RemoveRequestMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) IndependentConversationQueueing() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnection) EnableMultipleThreads() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RootProxy() (*NSDistantObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) DefaultConnection() (*NSConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) Statistics() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) CurrentConversation() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) InitWithReceivePort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RemoveRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) SetReplyTimeout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) SetIndependentConversationQueueing() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) MultipleThreadsEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RemoteObjects() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) AllConnections() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) AddRequestMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) RequestTimeout() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSConnection) SetRequestTimeout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) LocalObjects() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) IsValid() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RequestModes() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) ServiceConnectionWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) ConnectionWithReceivePort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) DispatchWithComponents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) SendPort() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RootProxyForConnectionWithRegisteredName() (*NSDistantObject, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) Invalidate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) RegisterName() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RunInNewThread() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) ReceivePort() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) AddRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSConnection) ReplyTimeout() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSConnection) RootObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnection) SetRootObject() error {
  return fmt.Errorf("unimplemented")
}


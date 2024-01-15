//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSNetService : objc.NSObject
*/

type NSNetService struct {
  *objc.NSObject

}

func (r *NSNetService) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) ScheduleInRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) GetInputStream() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Domain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) PublishWithOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Addresses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) InitWithDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) DictionaryFromTXTRecordData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) SetTXTRecordData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) RemoveFromRunLoop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Port() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) StartMonitoring() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Resolve() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) ResolveWithTimeout() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) StopMonitoring() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Publish() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) TXTRecordData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) SetIncludesPeerToPeer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Stop() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) DataFromTXTRecordDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) IncludesPeerToPeer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Type() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) HostName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


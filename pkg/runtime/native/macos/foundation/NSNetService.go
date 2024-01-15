//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSNetService : objc.NSObject
*/

type NSNetService struct {
  *objc.NSObject

}

func (r *NSNetService) InitWithDomain() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Resolve() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) HostName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Publish() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) DataFromTXTRecordDictionary() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Type() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) GetInputStream() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) IncludesPeerToPeer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Addresses() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) ScheduleInRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) PublishWithOptions() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) SetIncludesPeerToPeer() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) Port() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSNetService) Stop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) RemoveFromRunLoop() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) TXTRecordData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) StartMonitoring() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) DictionaryFromTXTRecordData() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSNetService) ResolveWithTimeout() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) SetTXTRecordData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSNetService) StopMonitoring() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSNetService) Domain() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


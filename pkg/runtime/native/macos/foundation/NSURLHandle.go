//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLHandle : objc.NSObject
*/

type NSURLHandle struct {
  *objc.NSObject

}

func (r *NSURLHandle) PropertyForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) WriteProperty() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) URLHandleClassForURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) LoadInBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) CancelLoadInBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) ResourceData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) AvailableResourceData() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) CanInitWithURL() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) BeginLoadInBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) FailureReason() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) RemoveClient() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) CachedHandleForURL() (*NSURLHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) RegisterURLHandleClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) AddClient() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) FlushCachedData() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) WriteData() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) EndLoadInBackground() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) Status() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) ExpectedResourceDataSize() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) BackgroundLoadDidFailWithReason() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) DidLoadBytes() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) PropertyForKeyIfAvailable() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLHandle) LoadInForeground() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}


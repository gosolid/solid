//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLProtocol : objc.NSObject
*/

type NSURLProtocol struct {
  *objc.NSObject

}

func (r *NSURLProtocol) PropertyForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) RemovePropertyForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) Request() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) CanonicalRequestForRequest() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) StartLoading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) SetProperty() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) UnregisterClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) CachedResponse() (*NSCachedURLResponse, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) InitWithRequest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) RegisterClass() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) StopLoading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) Client() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) CanInitWithRequest() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) RequestIsCacheEquivalent() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


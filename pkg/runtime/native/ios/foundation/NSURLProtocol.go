//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLProtocol : objc.NSObject
*/

type NSURLProtocol struct {
  *objc.NSObject

}

func (r *NSURLProtocol) Request() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) InitWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) StartLoading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) RemovePropertyForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) RegisterClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) Client() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) CachedResponse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) CanInitWithRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) CanonicalRequestForRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) PropertyForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) SetProperty() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) RequestIsCacheEquivalent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) StopLoading() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtocol) UnregisterClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


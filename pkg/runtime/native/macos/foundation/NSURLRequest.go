//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLRequest : objc.NSObject
*/

type NSURLRequest struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
  *NSMutableCopying
}

func (r *NSURLRequest) AllowsConstrainedNetworkAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AssumesHTTP3Capable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) Attribution() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) RequestWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AllowsCellularAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) CachePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) MainDocumentURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) RequiresDNSSECValidation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) SupportsSecureCoding() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) NetworkServiceType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AllowsExpensiveNetworkAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) URL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) TimeoutInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}


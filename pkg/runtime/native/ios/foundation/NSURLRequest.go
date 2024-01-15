//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLRequest : objc.NSObject
*/

type NSURLRequest struct {
  *objc.NSObject

}

func (r *NSURLRequest) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AllowsCellularAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AllowsExpensiveNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AllowsConstrainedNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) AssumesHTTP3Capable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) RequestWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) NetworkServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) RequiresDNSSECValidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) SupportsSecureCoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) CachePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) TimeoutInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) MainDocumentURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLRequest) Attribution() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


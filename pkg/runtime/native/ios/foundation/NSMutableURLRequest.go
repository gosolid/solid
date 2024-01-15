//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableURLRequest : Foundation.NSURLRequest
*/

type NSMutableURLRequest struct {
  *NSURLRequest

}

func (r *NSMutableURLRequest) SetCachePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) MainDocumentURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) NetworkServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAssumesHTTP3Capable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetRequiresDNSSECValidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetMainDocumentURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AllowsCellularAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAllowsCellularAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) Attribution() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) URL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AllowsExpensiveNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AssumesHTTP3Capable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAttribution() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) RequiresDNSSECValidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAllowsExpensiveNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AllowsConstrainedNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAllowsConstrainedNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) CachePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) TimeoutInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetTimeoutInterval() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetNetworkServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


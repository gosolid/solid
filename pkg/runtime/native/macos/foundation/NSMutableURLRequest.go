//js:package native/macos/foundation
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

func (r *NSMutableURLRequest) SetCachePolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAllowsCellularAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAssumesHTTP3Capable() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) Attribution() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAttribution() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) RequiresDNSSECValidation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) CachePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetNetworkServiceType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetTimeoutInterval() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) MainDocumentURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetMainDocumentURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AllowsCellularAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAllowsConstrainedNetworkAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AssumesHTTP3Capable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) TimeoutInterval() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AllowsExpensiveNetworkAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetAllowsExpensiveNetworkAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) AllowsConstrainedNetworkAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) SetRequiresDNSSECValidation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) URL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableURLRequest) NetworkServiceType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


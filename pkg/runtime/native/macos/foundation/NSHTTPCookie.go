//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSHTTPCookie : objc.NSObject
*/

type NSHTTPCookie struct {
  *objc.NSObject

}

func (r *NSHTTPCookie) CookieWithProperties() (*NSHTTPCookie, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) CookiesWithResponseHeaderFields() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Properties() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) IsSecure() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) CommentURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) InitWithProperties() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) IsHTTPOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Comment() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Version() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Value() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) ExpiresDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Path() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) PortList() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) SameSitePolicy() (**NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) RequestHeaderFieldsWithCookies() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) IsSessionOnly() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Domain() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


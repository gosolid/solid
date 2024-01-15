//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSHTTPCookie : objc.NSObject
*/

type NSHTTPCookie struct {
  *objc.NSObject

}

func (r *NSHTTPCookie) RequestHeaderFieldsWithCookies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) CookiesWithResponseHeaderFields() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Version() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) IsSessionOnly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Path() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) ExpiresDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) CommentURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) InitWithProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Properties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Domain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) IsHTTPOnly() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Comment() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) PortList() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) CookieWithProperties() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) Value() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) IsSecure() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookie) SameSitePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


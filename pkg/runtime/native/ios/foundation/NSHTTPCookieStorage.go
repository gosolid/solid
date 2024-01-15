//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSHTTPCookieStorage : objc.NSObject
*/

type NSHTTPCookieStorage struct {
  *objc.NSObject

}

func (r *NSHTTPCookieStorage) SharedCookieStorageForGroupContainerIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) DeleteCookie() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) RemoveCookiesSinceDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SetCookies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SortedCookiesUsingDescriptors() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SharedHTTPCookieStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SetCookieAcceptPolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SetCookie() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) CookiesForURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) Cookies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) CookieAcceptPolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


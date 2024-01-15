//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSHTTPCookieStorage : objc.NSObject
*/

type NSHTTPCookieStorage struct {
  *objc.NSObject

}

func (r *NSHTTPCookieStorage) Cookies() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SharedCookieStorageForGroupContainerIdentifier() (*NSHTTPCookieStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) RemoveCookiesSinceDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SetCookies() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SortedCookiesUsingDescriptors() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SharedHTTPCookieStorage() (*NSHTTPCookieStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SetCookie() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) DeleteCookie() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) CookiesForURL() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) CookieAcceptPolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSHTTPCookieStorage) SetCookieAcceptPolicy() error {
  return fmt.Errorf("unimplemented")
}


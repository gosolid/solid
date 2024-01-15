//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLProtectionSpace : objc.NSObject
*/

type NSURLProtectionSpace struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSURLProtectionSpace) IsProxy() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Host() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Port() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) AuthenticationMethod() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) InitWithHost() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) InitWithProxyHost() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) ReceivesCredentialSecurely() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Realm() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) ProxyType() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Protocol() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


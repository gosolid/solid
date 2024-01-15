//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLProtectionSpace : objc.NSObject
*/

type NSURLProtectionSpace struct {
  *objc.NSObject

}

func (r *NSURLProtectionSpace) Host() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) ProxyType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) AuthenticationMethod() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) InitWithHost() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Realm() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) ReceivesCredentialSecurely() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Protocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) InitWithProxyHost() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) IsProxy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLProtectionSpace) Port() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


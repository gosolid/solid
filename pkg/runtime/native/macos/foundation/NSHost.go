//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSHost : objc.NSObject
*/

type NSHost struct {
  *objc.NSObject

}

func (r *NSHost) CurrentHost() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) HostWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) IsHostCacheEnabled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHost) FlushHostCache() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHost) Names() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) Addresses() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) HostWithAddress() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) IsEqualToHost() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSHost) SetHostCacheEnabled() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSHost) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) Address() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHost) LocalizedName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSCachedURLResponse : objc.NSObject
*/

type NSCachedURLResponse struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSCachedURLResponse) InitWithResponse() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) Response() (*NSURLResponse, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) Data() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) UserInfo() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) StoragePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


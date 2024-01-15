//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSCachedURLResponse : objc.NSObject
*/

type NSCachedURLResponse struct {
  *objc.NSObject

}

func (r *NSCachedURLResponse) Data() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) UserInfo() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) StoragePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) InitWithResponse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCachedURLResponse) Response() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


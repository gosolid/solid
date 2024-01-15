//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLCache : objc.NSObject
*/

type NSURLCache struct {
  *objc.NSObject

}

func (r *NSURLCache) InitWithMemoryCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) StoreCachedResponse() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) RemoveCachedResponseForRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) RemoveAllCachedResponses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SharedURLCache() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) CurrentMemoryUsage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) CachedResponseForRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) RemoveCachedResponsesSinceDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SetSharedURLCache() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) DiskCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) CurrentDiskUsage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) MemoryCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SetMemoryCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SetDiskCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


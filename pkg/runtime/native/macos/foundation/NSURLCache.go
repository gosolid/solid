//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLCache : objc.NSObject
*/

type NSURLCache struct {
  *objc.NSObject

}

func (r *NSURLCache) RemoveCachedResponsesSinceDate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCache) InitWithMemoryCapacity() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) CachedResponseForRequest() (*NSCachedURLResponse, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SetMemoryCapacity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCache) DiskCapacity() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) RemoveCachedResponseForRequest() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SharedURLCache() (*NSURLCache, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SetSharedURLCache() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCache) MemoryCapacity() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) CurrentMemoryUsage() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) CurrentDiskUsage() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLCache) StoreCachedResponse() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCache) RemoveAllCachedResponses() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLCache) SetDiskCapacity() error {
  return fmt.Errorf("unimplemented")
}


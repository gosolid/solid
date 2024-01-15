//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLQueryItem : objc.NSObject
*/

type NSURLQueryItem struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSURLQueryItem) Value() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLQueryItem) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLQueryItem) QueryItemWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLQueryItem) Name() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


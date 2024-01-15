//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLQueryItem : objc.NSObject
*/

type NSURLQueryItem struct {
  *objc.NSObject

}

func (r *NSURLQueryItem) InitWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLQueryItem) QueryItemWithName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLQueryItem) Name() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLQueryItem) Value() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


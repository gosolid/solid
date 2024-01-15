//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSDictionary : objc.NSObject
*/

type NSDictionary struct {
  *objc.NSObject

}

func (r *NSDictionary) Count() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) KeyEnumerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) InitWithObjects() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDictionary) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


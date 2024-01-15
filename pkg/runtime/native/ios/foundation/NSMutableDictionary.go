//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableDictionary : Foundation.NSDictionary
*/

type NSMutableDictionary struct {
  *NSDictionary

}

func (r *NSMutableDictionary) RemoveObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableDictionary) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableDictionary) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableDictionary) InitWithCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableDictionary) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


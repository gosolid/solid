//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableSet : Foundation.NSSet
*/

type NSMutableSet struct {
  *NSSet

}

func (r *NSMutableSet) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableSet) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableSet) InitWithCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableSet) AddObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableSet) RemoveObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


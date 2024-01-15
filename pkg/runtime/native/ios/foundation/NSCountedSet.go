//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSCountedSet : Foundation.NSMutableSet
*/

type NSCountedSet struct {
  *NSMutableSet

}

func (r *NSCountedSet) InitWithSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) CountForObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) ObjectEnumerator() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) AddObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) RemoveObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) InitWithCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) InitWithArray() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


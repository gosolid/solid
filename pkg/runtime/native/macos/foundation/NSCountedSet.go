//js:package native/macos/foundation
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

func (r *NSCountedSet) RemoveObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) InitWithCapacity() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) InitWithArray() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) InitWithSet() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) CountForObject() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) ObjectEnumerator() (*NSEnumerator, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCountedSet) AddObject() error {
  return fmt.Errorf("unimplemented")
}


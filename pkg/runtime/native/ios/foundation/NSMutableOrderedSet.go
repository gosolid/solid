//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableOrderedSet : Foundation.NSOrderedSet
*/

type NSMutableOrderedSet struct {
  *NSOrderedSet

}

func (r *NSMutableOrderedSet) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableOrderedSet) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableOrderedSet) InitWithCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableOrderedSet) InsertObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableOrderedSet) RemoveObjectAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableOrderedSet) ReplaceObjectAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


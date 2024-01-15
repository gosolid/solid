//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableArray : Foundation.NSArray
*/

type NSMutableArray struct {
  *NSArray

}

func (r *NSMutableArray) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) AddObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) InsertObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) RemoveLastObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) RemoveObjectAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) ReplaceObjectAtIndex() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) InitWithCapacity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


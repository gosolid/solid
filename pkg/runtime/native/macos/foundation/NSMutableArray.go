//js:package native/macos/foundation
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

func (r *NSMutableArray) RemoveObjectAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) ReplaceObjectAtIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) InitWithCapacity() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) AddObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) InsertObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableArray) RemoveLastObject() error {
  return fmt.Errorf("unimplemented")
}


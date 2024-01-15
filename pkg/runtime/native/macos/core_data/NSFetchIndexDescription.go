//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSFetchIndexDescription : objc.NSObject
*/

type NSFetchIndexDescription struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
}

func (r *NSFetchIndexDescription) Entity() (*NSEntityDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) PartialIndexPredicate() (*foundation.NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) SetPartialIndexPredicate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) InitWithName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) Name() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) SetName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) Elements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexDescription) SetElements() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSFetchIndexElementDescription : objc.NSObject
*/

type NSFetchIndexElementDescription struct {
  *objc.NSObject
  *foundation.NSCoding
  *foundation.NSCopying
}

func (r *NSFetchIndexElementDescription) IsAscending() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) SetAscending() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) IndexDescription() (*NSFetchIndexDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) InitWithProperty() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) Property() (*NSPropertyDescription, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) PropertyName() (*foundation.NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) CollationType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSFetchIndexElementDescription) SetCollationType() error {
  return fmt.Errorf("unimplemented")
}


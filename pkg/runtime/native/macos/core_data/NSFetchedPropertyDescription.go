//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface CoreData.NSFetchedPropertyDescription : CoreData.NSPropertyDescription
*/

type NSFetchedPropertyDescription struct {
  *NSPropertyDescription

}

func (r *NSFetchedPropertyDescription) FetchRequest() (*NSFetchRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchedPropertyDescription) SetFetchRequest() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSCompositeAttributeDescription : CoreData.NSAttributeDescription
*/

type NSCompositeAttributeDescription struct {
  *NSAttributeDescription

}

func (r *NSCompositeAttributeDescription) Elements() (*foundation.NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompositeAttributeDescription) SetElements() error {
  return fmt.Errorf("unimplemented")
}


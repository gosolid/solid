//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSDerivedAttributeDescription : CoreData.NSAttributeDescription
*/

type NSDerivedAttributeDescription struct {
  *NSAttributeDescription

}

func (r *NSDerivedAttributeDescription) DerivationExpression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDerivedAttributeDescription) SetDerivationExpression() error {
  return fmt.Errorf("unimplemented")
}


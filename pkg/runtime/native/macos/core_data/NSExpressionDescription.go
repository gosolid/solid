//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
)

/*
interface CoreData.NSExpressionDescription : CoreData.NSPropertyDescription
*/

type NSExpressionDescription struct {
  *NSPropertyDescription

}

func (r *NSExpressionDescription) SetExpressionResultType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExpressionDescription) Expression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpressionDescription) SetExpression() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExpressionDescription) ExpressionResultType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


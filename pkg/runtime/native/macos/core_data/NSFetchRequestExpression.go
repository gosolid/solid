//js:package native/macos/core-data
package core_data

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/foundation"
  "fmt"
)

/*
interface CoreData.NSFetchRequestExpression : Foundation.NSExpression
*/

type NSFetchRequestExpression struct {
  *foundation.NSExpression

}

func (r *NSFetchRequestExpression) ExpressionForFetch() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequestExpression) RequestExpression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequestExpression) ContextExpression() (*foundation.NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSFetchRequestExpression) IsCountOnlyRequest() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}


//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSComparisonPredicate : Foundation.NSPredicate
*/

type NSComparisonPredicate struct {
  *NSPredicate

}

func (r *NSComparisonPredicate) PredicateOperatorType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) PredicateWithLeftExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) InitWithLeftExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) LeftExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) RightExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) CustomSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) Options() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) ComparisonPredicateModifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


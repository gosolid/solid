//js:package native/macos/foundation
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

func (r *NSComparisonPredicate) PredicateWithLeftExpression() (*NSComparisonPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) ComparisonPredicateModifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) LeftExpression() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) RightExpression() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) CustomSelector() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) Options() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) InitWithLeftExpression() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSComparisonPredicate) PredicateOperatorType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


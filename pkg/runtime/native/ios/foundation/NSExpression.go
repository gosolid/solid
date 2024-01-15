//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSExpression : objc.NSObject
*/

type NSExpression struct {
  *objc.NSObject

}

func (r *NSExpression) ExpressionForMinusSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) InitWithExpressionType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ConstantValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) LeftExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForEvaluatedObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForAggregate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForIntersectSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForConditional() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Operand() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) FalseExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionWithFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForVariable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionValueWithObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) RightExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForKeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForFunction() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForUnionSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForConstantValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForSubquery() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Collection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Predicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForAnyKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Function() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) TrueExpression() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Arguments() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) AllowEvaluation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) KeyPath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Variable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


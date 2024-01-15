//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSExpression : objc.NSObject
*/

type NSExpression struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSExpression) ExpressionWithFormat() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForVariable() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) InitWithExpressionType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) AllowEvaluation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSExpression) KeyPath() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Predicate() (*NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForMinusSet() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForAnyKey() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Collection() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) LeftExpression() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) TrueExpression() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) FalseExpression() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForFunction() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForUnionSet() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Operand() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForConstantValue() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Variable() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionBlock() (func(...any) (any, error), error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForKeyPath() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForConditional() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionValueWithObject() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ConstantValue() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForEvaluatedObject() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForBlock() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForSubquery() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Arguments() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) RightExpression() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForAggregate() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) ExpressionForIntersectSet() (*NSExpression, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSExpression) Function() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPredicate : objc.NSObject
*/

type NSPredicate struct {
  *objc.NSObject

}

func (r *NSPredicate) PredicateWithSubstitutionVariables() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) EvaluateWithObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) AllowEvaluation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateWithFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateFromMetadataQueryString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateWithValue() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateWithBlock() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


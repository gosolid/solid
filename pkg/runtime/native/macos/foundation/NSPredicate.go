//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSPredicate : objc.NSObject
*/

type NSPredicate struct {
  *objc.NSObject
  *NSSecureCoding
  *NSCopying
}

func (r *NSPredicate) PredicateWithValue() (*NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateWithBlock() (*NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateWithSubstitutionVariables() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) EvaluateWithObject() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) AllowEvaluation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateFormat() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateWithFormat() (*NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPredicate) PredicateFromMetadataQueryString() (*NSPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}


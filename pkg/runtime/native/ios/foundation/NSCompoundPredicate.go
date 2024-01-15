//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSCompoundPredicate : Foundation.NSPredicate
*/

type NSCompoundPredicate struct {
  *NSPredicate

}

func (r *NSCompoundPredicate) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) AndPredicateWithSubpredicates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) OrPredicateWithSubpredicates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) NotPredicateWithSubpredicate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) CompoundPredicateType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) Subpredicates() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) InitWithType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


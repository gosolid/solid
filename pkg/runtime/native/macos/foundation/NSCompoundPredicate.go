//js:package native/macos/foundation
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

func (r *NSCompoundPredicate) NotPredicateWithSubpredicate() (*NSCompoundPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) CompoundPredicateType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) Subpredicates() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) InitWithType() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) AndPredicateWithSubpredicates() (*NSCompoundPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCompoundPredicate) OrPredicateWithSubpredicates() (*NSCompoundPredicate, error) {
  return nil, fmt.Errorf("unimplemented")
}


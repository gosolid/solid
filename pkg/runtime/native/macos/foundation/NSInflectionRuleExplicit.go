//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSInflectionRuleExplicit : Foundation.NSInflectionRule
*/

type NSInflectionRuleExplicit struct {
  *NSInflectionRule

}

func (r *NSInflectionRuleExplicit) InitWithMorphology() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInflectionRuleExplicit) Morphology() (*NSMorphology, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/ios/foundation
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

func (r *NSInflectionRuleExplicit) Morphology() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInflectionRuleExplicit) InitWithMorphology() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


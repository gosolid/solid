//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMorphologyPronoun : objc.NSObject
*/

type NSMorphologyPronoun struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSMorphologyPronoun) DependentMorphology() (*NSMorphology, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) InitWithPronoun() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) Pronoun() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) Morphology() (*NSMorphology, error) {
  return nil, fmt.Errorf("unimplemented")
}


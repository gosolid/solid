//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMorphologyPronoun : objc.NSObject
*/

type NSMorphologyPronoun struct {
  *objc.NSObject

}

func (r *NSMorphologyPronoun) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) InitWithPronoun() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) Pronoun() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) Morphology() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphologyPronoun) DependentMorphology() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


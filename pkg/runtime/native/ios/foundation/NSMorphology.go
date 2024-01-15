//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSMorphology : objc.NSObject
*/

type NSMorphology struct {
  *objc.NSObject

}

func (r *NSMorphology) Number() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetNumber() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) PronounType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetPronounType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) GrammaticalGender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) GrammaticalCase() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetGrammaticalCase() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) Determination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) Definiteness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetPartOfSpeech() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) PartOfSpeech() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetDetermination() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) GrammaticalPerson() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetGrammaticalPerson() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetDefiniteness() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetGrammaticalGender() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


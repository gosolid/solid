//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSMorphology : objc.NSObject
*/

type NSMorphology struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSMorphology) SetNumber() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetDetermination() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetGrammaticalPerson() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetGrammaticalGender() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) PartOfSpeech() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetGrammaticalCase() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) PronounType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) GrammaticalGender() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetPartOfSpeech() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) Determination() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetPronounType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) Definiteness() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) SetDefiniteness() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMorphology) Number() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) GrammaticalCase() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSMorphology) GrammaticalPerson() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}


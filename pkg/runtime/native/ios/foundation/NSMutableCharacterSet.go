//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableCharacterSet : Foundation.NSCharacterSet
*/

type NSMutableCharacterSet struct {
  *NSCharacterSet

}

func (r *NSMutableCharacterSet) FormIntersectionWithCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) WhitespaceCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) DecimalDigitCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) IllegalCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) FormUnionWithCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) LowercaseLetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) PunctuationCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) AddCharactersInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) RemoveCharactersInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) AddCharactersInString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) ControlCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CapitalizedLetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithBitmapRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) RemoveCharactersInString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) Invert() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) LetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) UppercaseLetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) AlphanumericCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithCharactersInString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) WhitespaceAndNewlineCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) DecomposableCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) NewlineCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithContentsOfFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) NonBaseCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) SymbolCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


//js:package native/macos/foundation
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
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
}

func (r *NSMutableCharacterSet) AddCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) WhitespaceCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) DecimalDigitCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) LowercaseLetterCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) PunctuationCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) NewlineCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithRange() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) AddCharactersInString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) ControlCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithContentsOfFile() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) FormIntersectionWithCharacterSet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) Invert() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithCharactersInString() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) FormUnionWithCharacterSet() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) UppercaseLetterCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) DecomposableCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CapitalizedLetterCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) SymbolCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) AlphanumericCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) RemoveCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) RemoveCharactersInString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) WhitespaceAndNewlineCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) IllegalCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) CharacterSetWithBitmapRepresentation() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) LetterCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableCharacterSet) NonBaseCharacterSet() (*NSMutableCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}


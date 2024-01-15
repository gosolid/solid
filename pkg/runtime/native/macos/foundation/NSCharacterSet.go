//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSCharacterSet : objc.NSObject
*/

type NSCharacterSet struct {
  *objc.NSObject
  *NSCopying
  *NSMutableCopying
  *NSSecureCoding
}

func (r *NSCharacterSet) HasMemberInPlane() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) WhitespaceCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) UppercaseLetterCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) PunctuationCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) NewlineCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithCharactersInString() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterIsMember() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) ControlCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) WhitespaceAndNewlineCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) DecomposableCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) LongCharacterIsMember() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) DecimalDigitCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) IllegalCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CapitalizedLetterCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithContentsOfFile() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) LetterCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) NonBaseCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) BitmapRepresentation() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) InvertedSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithRange() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithBitmapRepresentation() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) IsSupersetOfSet() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) LowercaseLetterCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) AlphanumericCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) SymbolCharacterSet() (*NSCharacterSet, error) {
  return nil, fmt.Errorf("unimplemented")
}


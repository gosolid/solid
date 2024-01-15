//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSCharacterSet : objc.NSObject
*/

type NSCharacterSet struct {
  *objc.NSObject

}

func (r *NSCharacterSet) BitmapRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) WhitespaceAndNewlineCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterIsMember() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithCharactersInString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) InvertedSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) AlphanumericCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) UppercaseLetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) NewlineCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) LowercaseLetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) LongCharacterIsMember() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) HasMemberInPlane() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) WhitespaceCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) NonBaseCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) InitWithCoder() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) ControlCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) LetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) IllegalCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithContentsOfFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) IsSupersetOfSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) SymbolCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) DecimalDigitCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) DecomposableCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) PunctuationCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CapitalizedLetterCharacterSet() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSCharacterSet) CharacterSetWithBitmapRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


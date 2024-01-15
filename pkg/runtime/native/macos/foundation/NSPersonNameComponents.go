//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSPersonNameComponents : objc.NSObject
*/

type NSPersonNameComponents struct {
  *objc.NSObject
  *NSCopying
  *NSSecureCoding
}

func (r *NSPersonNameComponents) GivenName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) NameSuffix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetNamePrefix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetMiddleName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) Nickname() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetPhoneticRepresentation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetGivenName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) FamilyName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetFamilyName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetNameSuffix() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetNickname() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) PhoneticRepresentation() (*NSPersonNameComponents, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) NamePrefix() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) MiddleName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}


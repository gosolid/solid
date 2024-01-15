//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSPersonNameComponents : objc.NSObject
*/

type NSPersonNameComponents struct {
  *objc.NSObject

}

func (r *NSPersonNameComponents) NamePrefix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) GivenName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetFamilyName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetNamePrefix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) FamilyName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetNickname() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetGivenName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) Nickname() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) PhoneticRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetPhoneticRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) MiddleName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetMiddleName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) NameSuffix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPersonNameComponents) SetNameSuffix() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


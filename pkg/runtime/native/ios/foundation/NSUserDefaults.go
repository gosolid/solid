//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
)

/*
interface Foundation.NSUserDefaults : objc.NSObject
*/

type NSUserDefaults struct {
  *objc.NSObject

}

func (r *NSUserDefaults) StringForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DictionaryForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) BoolForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemoveVolatileDomainForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) PersistentDomainNames() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetFloat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RegisterDefaults() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) InitWithSuiteName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetBool() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetPersistentDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemovePersistentDomainForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ObjectIsForcedForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ResetStandardUserDefaults() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) InitWithUser() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DataForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) StringArrayForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) FloatForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DoubleForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetDouble() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemoveSuiteNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetVolatileDomain() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) URLForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) AddSuiteNamed() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemoveObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ArrayForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) IntegerForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetInteger() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) VolatileDomainForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) PersistentDomainForName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) Synchronize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DictionaryRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) StandardUserDefaults() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) VolatileDomainNames() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}


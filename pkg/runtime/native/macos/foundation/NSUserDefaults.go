//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUserDefaults : objc.NSObject
*/

type NSUserDefaults struct {
  *objc.NSObject

}

func (r *NSUserDefaults) URLForKey() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetInteger() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) VolatileDomainForName() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) PersistentDomainForName() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) StandardUserDefaults() (*NSUserDefaults, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemoveObjectForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ArrayForKey() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) FloatForKey() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetBool() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RegisterDefaults() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DictionaryRepresentation() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemovePersistentDomainForName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ObjectIsForcedForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ResetStandardUserDefaults() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) InitWithUser() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DictionaryForKey() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) StringArrayForKey() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DoubleForKey() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetVolatileDomain() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetPersistentDomain() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) VolatileDomainNames() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) IntegerForKey() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) DataForKey() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) BoolForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemoveSuiteNamed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) RemoveVolatileDomainForName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) InitWithSuiteName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) StringForKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetFloat() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetDouble() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) Synchronize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) SetURL() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) AddSuiteNamed() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserDefaults) PersistentDomainNames() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}


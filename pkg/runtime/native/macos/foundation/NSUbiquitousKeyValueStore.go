//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUbiquitousKeyValueStore : objc.NSObject
*/

type NSUbiquitousKeyValueStore struct {
  *objc.NSObject

}

func (r *NSUbiquitousKeyValueStore) SetArray() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetBool() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) Synchronize() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) StringForKey() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) ArrayForKey() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DictionaryForKey() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetString() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) LongLongForKey() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) BoolForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetDouble() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DefaultStore() (*NSUbiquitousKeyValueStore, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) ObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) RemoveObjectForKey() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DataForKey() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DoubleForKey() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetLongLong() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DictionaryRepresentation() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetData() error {
  return fmt.Errorf("unimplemented")
}


//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSUbiquitousKeyValueStore : objc.NSObject
*/

type NSUbiquitousKeyValueStore struct {
  *objc.NSObject

}

func (r *NSUbiquitousKeyValueStore) ArrayForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DictionaryForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) LongLongForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DoubleForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetArray() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DictionaryRepresentation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) StringForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DataForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetLongLong() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) DefaultStore() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) ObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) RemoveObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetDouble() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) Synchronize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) BoolForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUbiquitousKeyValueStore) SetBool() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

